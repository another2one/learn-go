package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

func Test_match(t *testing.T) {
	mBody, err := ioutil.ReadFile("fail.html")
	if err != nil {
		t.Fatal(err)
	}

	// 匹配歌曲链接
	reg1 := regexp.MustCompile(`(https\:\/\/apis\.jxcxin\.cn.*?)\'`)
	urlRes1 := reg1.FindSubmatch(mBody)
	if len(urlRes1) == 0 {
		t.Fatal("not match download url:")
	}
	dUrl := strings.Replace(string(urlRes1[1]), "amp;", "", -1)
	fmt.Printf("nusic url: %q", dUrl)

	// 匹配歌曲类型
	reg3 := regexp.MustCompile(`type=(.*)`)
	urlRes3 := reg3.FindSubmatch([]byte(dUrl))
	if len(urlRes3) == 0 {
		t.Fatal("not match music type:")
	}
	musicType := string(urlRes3[1])

	// 匹配歌曲名
	reg2 := regexp.MustCompile(`<title>(.*?)-`)
	urlRes2 := reg2.FindSubmatch(mBody)
	if len(urlRes2) == 0 {
		t.Fatal("not match mp3 name:")
	}
	musicName := string(urlRes2[1])

	// 保存路径
	filePath := fmt.Sprintf("download/%s-周杰伦.%s", musicName, musicType)
	fmt.Printf("music path %q \n", filePath)
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("not found:", err)
	} else {
		fmt.Printf("%v \n", info)
	}
}
