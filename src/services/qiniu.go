package services

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"time"
)

var QiniuClient *Qiniu

type Qiniu struct {
	AK     string
	SK     string
	Bucket string
}

func NewQiniu() *Qiniu {
	if QiniuClient != nil {
		return QiniuClient
	}
	ak := viper.GetString("QINIU_AK")
	sk := viper.GetString("QINIU_SK")
	bucket := viper.GetString("QINIU_BUCKET")
	QiniuClient := &Qiniu{
		AK:     ak,
		SK:     sk,
		Bucket: bucket,
	}
	return QiniuClient
}

func (q Qiniu) Upload(path string, key string) (string, error) {
	token, err := q.getAccessToken()
	if err != nil {
		return token, err
	}
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	err = formUploader.PutFile(context.Background(), &ret, token, key, path, &putExtra)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash)
	return ret.Key, err
}
func (q Qiniu) getAccessToken() (string, error) {
	cacheKey := "qiniu_toen"
	info, err := NewRedisClient().Get(cacheKey).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	if info != "" {
		return info, nil
	}
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(q.AK, q.SK)
	upToken := putPolicy.UploadToken(mac)
	if upToken != "" {
		_ = NewRedisClient().Set(cacheKey, upToken, time.Second*7000).Err()
	}
	return upToken, nil
}
