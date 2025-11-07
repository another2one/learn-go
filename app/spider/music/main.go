package main

import (
	"errors"
	"fmt"
	"io"
	"learn-go/common/funcs"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	searchUrl    = "https://www.gequbao.com/s/"
	host         = "https://www.gequbao.com/"
	keyWord      string
	ErrorFileOp  = errors.New("op file error")
	ErrorUrlGet  = errors.New("url get error")
	downLoadPath = funcs.ProjectPath + "runtime/download/"
)

func main() {
	var body []byte
	var err error
	//fmt.Println("https://www.gequbao.com/ 歌曲爬虫，请输入关键字")
	//for {
	//	fmt.Scanln(&keyWord)
	//	if len(keyWord) == 0 {
	//		fmt.Printf("输入有误 %s \n", keyWord)
	//		continue
	//	}
	//	//获取搜索结果
	//	body, err = get(searchUrl + keyWord)
	//	if err != nil {
	//		fmt.Printf("search error: %s \n", err)
	//		continue
	//	}
	//	break
	//}
	//获取搜索结果
	keyWord = "孙燕姿"
	body, err = get(searchUrl + keyWord)
	if err != nil {
		log.Fatalf("search error: %s \n", err)
	}
	// 匹配歌曲链接
	reg := regexp.MustCompile(`href=\"\/(music\/\w+)\"`)
	urlRes := reg.FindAllSubmatch(body, -1)
	lens := len(urlRes)
	if lens == 0 {
		log.Fatalln("not match url:", body)
	}
	urls := make([]string, 0)
	urlMap := make(map[string]struct{}, lens)
	for _, v := range urlRes {
		url := strings.Trim(string(v[1]), " ")
		if _, ok := urlMap[url]; !ok {
			urlMap[url] = struct{}{}
			urls = append(urls, url)
		}
	}
	fmt.Printf("%q \n", urls)

	// 详情页
	for _, v1 := range urls {
		if len(v1) > 0 {
			id, ok := strings.CutPrefix(v1, "music/")
			if !ok {
				fmt.Errorf("url error %s \n", v1)
				continue
			}
			fmt.Printf("%v \n", id)
			detailPage(host+v1, id)
			break
			// 休眠一会 防止被检测出来爬虫
			time.Sleep(time.Millisecond * 1500)
		}
	}
}

// detailPage 下载歌曲
// id 歌曲在爬虫网站的id 用于防止重复下载
func detailPage(url string, id string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	mBody, err := get(url)
	if err != nil {
		panic(err)
	}

	// 匹配歌曲链接
	reg1 := regexp.MustCompile(`(https\:\/\/apis\.jxcxin\.cn.*?)\'`)
	urlRes1 := reg1.FindSubmatch(mBody)
	if len(urlRes1) == 0 {
		if err = os.WriteFile("fail.html", mBody, 0666); err != nil {
			panic(fmt.Errorf("write file error: %s [%s]", "fail.html", err))
		}
		panic("music url error")
	}
	dUrl := strings.Replace(string(urlRes1[1]), "amp;", "", -1)
	fmt.Printf("music url: %q \n", dUrl)

	// 匹配歌曲类型
	reg3 := regexp.MustCompile(`type=(.*)`)
	urlRes3 := reg3.FindSubmatch([]byte(dUrl))
	if len(urlRes3) == 0 {
		if err = os.WriteFile("fail.html", mBody, 0666); err != nil {
			panic(fmt.Errorf("write file error: %s [%s]", "fail.html", err))
		}
		panic("music type error")
	}
	musicType := string(urlRes3[1])

	// 匹配歌曲名
	reg2 := regexp.MustCompile(`<title>(.*?)-`)
	urlRes2 := reg2.FindSubmatch(mBody)
	if len(urlRes2) == 0 {
		if err = os.WriteFile("fail.html", mBody, 0666); err != nil {
			panic(fmt.Errorf("write file error: %s [%s]", "fail.html", err))
		}
		panic("music name error")
	}
	musicName := string(urlRes2[1])

	// 保存路径 歌曲名加字符防止网易云盘判断为vip歌曲
	filePath := fmt.Sprintf(downLoadPath+"%s-%s.%s", musicName, id, musicType)
	fmt.Printf("music path %q \n", filePath)
	_, err = os.Stat(filePath)
	if err == nil {
		fmt.Printf("%s has exists ... \n", filePath)
		return
	}

	// 下载歌曲
	mp3Body, err := get(dUrl)
	if err != nil {
		panic(fmt.Errorf("download error: %s [%s]", dUrl, err))
	}
	if err = os.WriteFile(filePath, mp3Body, 0666); err != nil {
		panic(fmt.Errorf("write file error: %s [%s]", filePath, err))
	}
	return
}

// get 获取 url 网页内容
func get(url string) ([]byte, error) {
	client := http.Client{}
	defer client.CloseIdleConnections()

	//获取搜索结果
	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	return io.ReadAll(rsp.Body)
}

func _log(params any) {
	fmt.Printf("%q", params)
}
