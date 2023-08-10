package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	searchUrl   = "https://www.gequbao.com/s/%E5%91%A8%E6%9D%B0%E4%BC%A6-"
	host        = "https://www.gequbao.com/"
	ErrorFileOp = errors.New("op file error")
	ErrorUrlGet = errors.New("url get error")
)

func main() {
	//获取搜索结果
	body, err := get(searchUrl)
	if err != nil {
		log.Fatalln("search error: ", err)
	}

	// 匹配歌曲链接
	// reg := regexp.MustCompile(`href=\"\/(music\/\w+)\"`)
	reg := regexp.MustCompil
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
	fmt.Printf("%q", urls)
	// 详情页
	for k, v1 := range urls {
		if len(v1) > 0 && k > 70 {
			detailPage(host + v1)
			// 休眠一会
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

func detailPage(url string) {
	defer func() {
		if err := recover(); err != nil {
			//fmt.Println(err)
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
		if err = ioutil.WriteFile("fail.html", mBody, 0666); err != nil {
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
		if err = ioutil.WriteFile("fail.html", mBody, 0666); err != nil {
			panic(fmt.Errorf("write file error: %s [%s]", "fail.html", err))
		}
		panic("music type error")
	}
	musicType := string(urlRes3[1])

	// 匹配歌曲名
	reg2 := regexp.MustCompile(`<title>(.*?)-`)
	urlRes2 := reg2.FindSubmatch(mBody)
	if len(urlRes2) == 0 {
		if err = ioutil.WriteFile("fail.html", mBody, 0666); err != nil {
			panic(fmt.Errorf("write file error: %s [%s]", "fail.html", err))
		}
		panic("music name error")
	}
	musicName := string(urlRes2[1])

	// 保存路径
	filePath := fmt.Sprintf("download/%s-周杰伦.%s", musicName, musicType)
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
	if err = ioutil.WriteFile(filePath, mp3Body, 0666); err != nil {
		panic(fmt.Errorf("write file error: %s [%s]", filePath, err))
	}
	return
}

func get(url string) ([]byte, error) {
	client := http.Client{}
	defer client.CloseIdleConnections()

	//获取搜索结果
	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	return ioutil.ReadAll(rsp.Body)
}

func _log(params any) {
	fmt.Printf("%q", params)
}
