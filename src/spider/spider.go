package spider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
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
	Url string
	Id  uint64
}

type WeiboAjax struct {
	Url string
	Id  uint64
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
		err := os.Mkdir(folderPath, os.ModePerm) //0777也可以os.ModePerm
		return err
	}
	return nil
}

func (p Weibo) Start() uint32 {
	htmlInfo := p.GetHtml(1)
	_ = p.ParseHtml(htmlInfo)
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
	Msgtype string `json:"msgtype"`
	Text    string `json:"text"`
}

func (postData *JsonPostSample) SamplePost() {
	bytesData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "https://oapi.dingtalk.com/robot/send?access_token=c725a8ecd243517e89fce48e5ae7d8d0d22feb5d7e52de22a515cb2656978c91"
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
