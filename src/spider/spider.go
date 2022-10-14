package spider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/spf13/viper"
	"golang.org/x/net/publicsuffix"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type Spider interface {
	// 方法列表
	Start() uint32
	GetHtml(page uint32) io.Reader
	ParseHtml(content io.Reader) int
}
type Joker struct {
	Url      string
	PageSize uint32
	Page     uint32
	EndPage  uint32
	CoreNum  uint32
}

type Weibo struct {
	Url     string
	Id      uint64
	ShotUrl string
	Path    string
}
type WeiboResult struct {
	Data struct {
		SinceId string `json:"since_id"`
		List    []struct {
			Visible struct {
				Type   int `json:"type"`
				ListId int `json:"list_id"`
			} `json:"visible"`
			CreatedAt string `json:"created_at"`
			Id        int64  `json:"id"`
			Idstr     string `json:"idstr"`
			Mid       string `json:"mid"`
			Mblogid   string `json:"mblogid"`
			User      struct {
				Id              int64  `json:"id"`
				Idstr           string `json:"idstr"`
				PcNew           int    `json:"pc_new"`
				ScreenName      string `json:"screen_name"`
				ProfileImageUrl string `json:"profile_image_url"`
				ProfileUrl      string `json:"profile_url"`
				Verified        bool   `json:"verified"`
				VerifiedType    int    `json:"verified_type"`
				Domain          string `json:"domain"`
				Weihao          string `json:"weihao"`
				VerifiedTypeExt int    `json:"verified_type_ext"`
				AvatarLarge     string `json:"avatar_large"`
				AvatarHd        string `json:"avatar_hd"`
				FollowMe        bool   `json:"follow_me"`
				Following       bool   `json:"following"`
				Mbrank          int    `json:"mbrank"`
				Mbtype          int    `json:"mbtype"`
				PlanetVideo     bool   `json:"planet_video"`
				IconList        []struct {
					Type string `json:"type"`
					Data struct {
						Mbrank int `json:"mbrank"`
						Mbtype int `json:"mbtype"`
					} `json:"data"`
				} `json:"icon_list"`
			} `json:"user"`
			CanEdit         bool   `json:"can_edit"`
			TextRaw         string `json:"text_raw"`
			Text            string `json:"text"`
			TextLength      int    `json:"textLength"`
			Source          string `json:"source"`
			Favorited       bool   `json:"favorited"`
			RepostsCount    int    `json:"reposts_count"`
			CommentsCount   int    `json:"comments_count"`
			AttitudesCount  int    `json:"attitudes_count"`
			AttitudesStatus int    `json:"attitudes_status"`
			IsLongText      bool   `json:"isLongText"`
			ShareRepostType int    `json:"share_repost_type"`
			IsTop           int    `json:"isTop"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
	Ok int `json:"ok"`
}

func (p Joker) Start() uint32 {
	err := CreateDateDir("./img")
	if err != nil {
		panic(err)
	}
	for i := p.Page; i <= p.EndPage; i++ {
		randNum := rand.Intn(100)
		time.Sleep(time.Duration(time.Duration(randNum) * time.Microsecond))
		htmlInfo := p.GetHtml(i)
		length := p.ParseHtml(htmlInfo)
		if length <= 0 {
			break
		}
	}
	return 1
}

func (p Joker) GetHtml(page uint32) io.Reader {
	//proxyURL, _ := url.Parse("http://127.0.0.1:8888")
	//trans := &http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}go bu
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	client := &http.Client{Timeout: 3 * time.Second, Jar: jar}
	requestUrl := fmt.Sprintf(p.Url, page)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	request.Header.Add("Host", "tieba.baidu.com")
	request.Header.Add("Proxy-Connection", "keep-alive")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	result, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	if result.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", result.StatusCode, result.Status)
	}
	return result.Body
}

func (p Joker) ParseHtml(content io.Reader) int {
	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		log.Fatal(err)
	}
	allImgUrl := make([]string, 0)
	doc.Find(".p_author_face ").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		img, _ := s.Html()
		r, _ := regexp.Compile(`data-tb-lazyload="(.*)"`)
		imgUrl := r.FindStringSubmatch(img)
		if len(imgUrl) > 1 {
			allImgUrl = append(allImgUrl, imgUrl[1])
		}
	})
	c := make(chan bool, 3)
	for _, v := range allImgUrl {
		go func(v string, c chan bool) {
			download(v, c)
		}(v, c)
	}
	i := 0
	for range c {
		i = i + 1
		if len(allImgUrl) <= i {
			close(c)
		}
	}
	return len(allImgUrl)
}

func download(path string, c chan bool) {
	if strings.HasPrefix(path, "//") {
		path = "http:" + path
	}
	res, err := http.Get(path)
	if err != nil {
		panic(err)
	}
	re3, _ := regexp.Compile("\\?t=.*")
	//把匹配的所有字符a替换成b

	_, fileName := filepath.Split(path)
	rep2 := re3.ReplaceAllString(fileName, "")
	f, err := os.Create("./imga/" + rep2 + ".jpg")
	if err != nil {
		panic(err)
	}
	_, _ = io.Copy(f, res.Body)
	fmt.Println(fileName)
	c <- true
}

func CreateDateDir(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		err := os.MkdirAll(folderPath, os.ModePerm) //0777也可以os.ModePerm
		return err
	}
	return nil
}

func (p Weibo) Start() uint32 {
	htmlInfo := p.GetHtml(1)
	//_ = p.ParseHtml(htmlInfo)
	_ = p.ParseJson(htmlInfo)
	return 1
}

func (p Weibo) GetHtml(page uint32) io.Reader {
	//proxyURL, _ := url.Parse("http://127.0.0.1:8888")
	//trans := &http.Transport{
	//	Proxy: http.ProxyURL(proxyURL),
	//}
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	client := &http.Client{Timeout: 30 * time.Second, Jar: jar}
	requestUrl := fmt.Sprintf(p.Url)
	fmt.Println(requestUrl)
	request, err := http.NewRequest("GET", requestUrl, nil)
	request.Header.Add("Host", "weibo.com")
	request.Header.Add("Proxy-Connection", "keep-alive")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("Sec-Fetch-Mode", "navigate")
	request.Header.Add("Sec-Fetch-User", "?1")
	request.Header.Add("Sec-Fetch-Mode", "navigate")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Cache-Control", "zh-CN,zh;q=0.9")
	request.Header.Add("Cookie", "SINAGLOBAL=9174360902242.89.1570602938922; SUB=_2AkMqoLFMf8NxqwJRmfwQxGLhaYV0yQ7EieKc_ECXJRMxHRl-yT9jql4ntRB6ASCfo592D3Ma02sMrAluN9K3ULu1SUUM; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9WhqfJQjkvIF1JOCAHqmWI4S; UOR=xue.t.eoffcn.com,widget.weibo.com,www.eoffcn.com; _s_tentry=-; Apache=5386194827390.069.1577437591802; ULV=1577437591827:21:9:4:5386194827390.069.1577437591802:1577348576677; Ugrow-G0=1ac418838b431e81ff2d99457147068c; TC-V5-G0=799b73639653e51a6d82fb007f218b2f; TC-Page-G0=c9fb286cd873ae77f97ce98d19abfb61|1577437646|1577437591")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")
	result, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	if result.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", result.StatusCode, result.Status)
	}
	return result.Body
}

func (p Weibo) ParseJson(content io.Reader) int {
	weiboResult := WeiboResult{}
	info, _ := ioutil.ReadAll(content)
	err := json.Unmarshal(info, &weiboResult)
	if err != nil {
		fmt.Println("content:", string(info))
		fmt.Println("err:", err)
		return 0
	}
	if weiboResult.Ok != 1 {

		fmt.Println("err ok:", weiboResult)
		return 0
	}
	type cacheData struct {
		Id        int64   `json:"id"`
		Mid       float64 `json:"mid"`
		Content   string  `json:"content"`
		Username  string  `json:"username"`
		CreatedAt string  `json:"created_at"`
		Path      string  `json:"path"`
	}
	allMid := make([]cacheData, 0)
	for _, row := range weiboResult.Data.List {
		rowMid, _ := strconv.ParseFloat(row.Mid, 64)

		tim, _ := time.Parse(time.RubyDate, row.CreatedAt)
		allMid = append(allMid, cacheData{
			Id:        row.User.Id,
			Mid:       rowMid,
			Content:   row.TextRaw,
			Username:  row.User.ScreenName,
			CreatedAt: tim.Format("2006-01-02 15:04:05"),
		})
	}

	if len(allMid) > 0 {
		sort.SliceStable(allMid, func(i, j int) bool {
			return allMid[i].Mid < allMid[j].Mid
		})

		dir := p.Path
		err = CreateDateDir(dir)
		if err != nil {
			panic(err)
		}
		latestMid := allMid[len(allMid)-1]
		fileName := strconv.FormatInt(int64(p.Id), 10)
		filePath := dir + "/" + fileName + ".txt"
		var f *os.File
		var con []byte
		maxMid := cacheData{}
		if isExist(filePath) {
			con, err = os.ReadFile(filePath)
			f, err = os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
			json.Unmarshal(con, &maxMid)
		} else {
			f, err = os.Create(filePath)
		}

		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
			return 0
		} else {

			if latestMid.Mid <= maxMid.Mid {
				latestMid = maxMid
			} else {
				fmt.Println(maxMid.Mid)
				fmt.Println(latestMid.Content)
				fmt.Println(latestMid.Username)
				fmt.Println("send")

				//截个图
				shotDir := time.Now().Format("20060102")
				dir := p.Path + "/shot/" + shotDir + "/"
				err := CreateDateDir(dir)
				if err != nil {
					fmt.Println("err:", err)
					return 0
				}
				shotFileName := fmt.Sprintf("/%d-%s.png", p.Id, time.Now().Format("20060102150405"))
				path := dir + shotFileName
				p.ScreenShot(path, p.ShotUrl)
				latestMid.Path = shotFileName
				fileUrl := viper.GetString("fileBaseUrl") + "/" + shotDir + shotFileName
				msg := "### " + latestMid.Username + "\n>" + latestMid.Content + "\n![screenshot](" + fileUrl + ")"
				postInfo := JsonPostSample{
					Msgtype: "markdown",
					Markdown: Markdown{
						Title: latestMid.Username,
						Text:  msg,
					},
				}
				postInfo.SamplePost()
			}
			cacheInfo, err := json.Marshal(latestMid)
			if err != nil {
				fmt.Println("cacheInfoErr:", err)
				return 0
			}
			fmt.Println(string(cacheInfo))
			_, err = f.Write(cacheInfo)
			if err != nil {
				fmt.Println("WriteErr:", err)
				return 0
			}
		}

	}

	return 1
}

func (p Weibo) ParseHtml(content io.Reader) int {
	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		log.Fatal(err)
	}
	html, _ := doc.Html()
	r, _ := regexp.Compile(`mid=\\"(\d+)\\"`)
	mids := r.FindAllStringSubmatch(html, -1)
	allMid := make([]float64, 0)
	for _, midInfo := range mids {
		curMid, _ := strconv.Atoi(midInfo[1])
		allMid = append(allMid, float64(curMid))
	}
	//<h1 class=\"username\">李子柒<\/h1>
	usernameR, _ := regexp.Compile(`<h1 class=\\"username\\">(.*)<\\/h1>`)
	username := usernameR.FindStringSubmatch(html)
	fmt.Println("allMid:", allMid)
	if len(allMid) > 0 {
		sort.Float64s(allMid)
		dir := "./cache"
		err = CreateDateDir(dir)
		if err != nil {
			panic(err)
		}
		latestMid := allMid[len(allMid)-1]
		fileName := strconv.FormatInt(int64(p.Id), 10)
		filePath := dir + "/" + fileName + ".txt"
		f, err := os.Open(filePath)
		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			con, err := ioutil.ReadAll(f)
			fmt.Println(string(con))
			fmt.Println(err)

			maxMid, _ := strconv.ParseFloat(string(con), 64)

			if latestMid <= maxMid {
				latestMid = maxMid
			} else {
				fmt.Println(latestMid)
				fmt.Println(maxMid)
				fmt.Println(username)
				fmt.Println("send")
				//postInfo := JsonPostSample{
				//	Msgtype:"text",
				//	Text: "{\"content\":\""+username[1]+"发微博了\"}",
				//}
				//postInfo.SamplePost()
			}
			_, err = f.Write([]byte(strconv.FormatFloat(latestMid, 'f', 4, 64)))

		}

	}

	return len(allMid)
}

type JsonPostSample struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (postData *JsonPostSample) SamplePost() {
	bytesData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := viper.GetString("DingUrl")
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func (p Weibo) ScreenShot(path string, url string) int {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(url, 80, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(path, buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
	return 1
}

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Emulate(device.IPhone6),
		chromedp.WaitVisible(`div[class="weibo-text"]`),
		chromedp.FullScreenshot(res, quality),
	}
}
