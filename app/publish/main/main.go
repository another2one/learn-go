package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"learn-go/common/funcs"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

var (
	path       string
	sleep      int
	limit      int
	cacheSlice []string
	sep        = " ++ "
	env        = "prod"
	host       = ""
	cacheFile  = "D:/" + env + ".txt"
)

type Article struct {
	Title    string `json:"title"`
	KeyWords string `json:"keywords"`
	Tags     string `json:"tags"`
	Content  string `json:"content"`
	Type     string `json:"type"`
	Desc     string `json:"desc"`
}

// 文章自动发布
// go run .\main.go -limit 3 -path D:/tmp -sleep 5
//  - limit 发多少篇文章
//  - path 文章位置
//  - sleep 每篇间隔时间 （秒）
func main() {

	flag.StringVar(&path, "path", "./", ``)
	flag.IntVar(&sleep, "sleep", 3, "inter time")
	flag.IntVar(&limit, "limit", 100000, "limit article pub num")
	flag.Parse()

	host = "http://wp.cc"
	if env == "prod" {
		host = "https://www.lekuaisong.com"
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	successNum := 0

	// 记录已发送文章
	file, err := os.OpenFile(cacheFile, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err : ", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	cacheSlice = make([]string, 10)
	str := readFile(cacheFile)
	if str != "" {
		cacheSlice = strings.Split(str, sep)
	}

	// 执行操作
	go func() {
		// 打开目录
		dir, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatalln("dir error")
		}
		// 遍历目录
	label1:
		for _, fi := range dir {
			if !fi.IsDir() {
				continue
			}

			if successNum >= limit {
				break label1
			}
			if fi.IsDir() && !in_array(fi.Name(), cacheSlice) {
				res := pubArticle(fi.Name())
				// 发布成功就记录
				if res {
					writer.WriteString(sep + fi.Name())
					successNum++
				}
				// 是否需要休息几秒
				if sleep > 0 {
					time.Sleep(time.Second * time.Duration(sleep))
				}
			} else {
				fmt.Printf("%s 已添加...... \n", fi.Name())
			}
		}
		fmt.Printf("%d 发布完成...... \n", successNum)
		writer.Flush()
	}()

	// 监听程序退出信号 将文件内容刷入磁盘
	s := <-c
	fmt.Println("Got signal:", s)
	writer.Flush()
}

// pubArticle 发布文章
func pubArticle(articleName string) bool {
	articlePath := path + "/" + articleName

	article := Article{
		Title:    articleName,
		KeyWords: readFile(articlePath + "/keywords.txt"),
		Content:  readFile(articlePath + "/" + articleName + ".txt"),
		Desc:     readFile(articlePath + "/description.txt"),
		Tags:     readFile(articlePath + "/tags.txt"),
		Type:     "",
	}
	if article.Content == "" {
		fmt.Printf("%s 的文章内容为空 ...... \n", articleName)
		return false
	}

	return postArticle(article)
}

// readFile 读取文件
func readFile(path string) string {
	str, err := ioutil.ReadFile(path)
	if err == nil {
		return string(str)
	}
	return ""
}

// in_array item 是否在切片 s 里面
func in_array(item string, s []string) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

// postArticle 执行文章发布
func postArticle(article Article) bool {
	aJson, err := json.Marshal(article)
	if err != nil {
		fmt.Println("encode error")
		return false
	}
	dateNow := strconv.FormatInt(time.Now().Unix(), 10)
	sign := funcs.Md5V(funcs.Md5V(string(aJson)+dateNow) + "sae32242esxae23rsadq2zsare234rdsdxc23rdfw23rcw34r2a165wdwdxw3e")

	postValue := url.Values{
		"sign":    {sign},
		"time":    {dateNow},
		"article": {string(aJson)},
	}

	client := http.Client{Timeout: time.Second * 3}
	// TODO: php 为啥接收不到参数
	// resp, err := http.PostForm("http://wp.cc/post.php", strings.NewReader("sign="+sign+"&&time="+dateNow+"&&article="+string(aJson)))
	// 这种会转义参数 接收端需要注意 ??? 为啥换成client就不需要转义
	resp, err := client.PostForm(host+"/post.php", postValue)
	if err != nil {
		fmt.Printf("%s 创建失败 ...... \n", article.Title)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s 获取返回数据失败 ...... \n", article.Title)
	}

	fmt.Printf("%s 	%s\n", article.Title, string(body))
	return true
}
