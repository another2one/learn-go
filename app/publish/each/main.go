package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 循环直接数据库写入文章

var (
	path  string
	sleep int
	limit int
	dbs   []*gorm.DB
)

type Article struct {
	Title    string `json:"title"`
	KeyWords string `json:"seo_keywords"`
	// Tags     string `json:"tags"`
	Content string `json:"content"`
	NodeID  int    `json:"node_id"` // 分类
	Desc    string `json:"seo_description"`
}

// 文章自动发布
// go run .\main.go -limit 3 -path D:/tmp -sleep 5
//  - limit 发多少篇文章
//  - path 文章位置
//  - sleep 每篇间隔时间 （秒）
// db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)
// db.Exec("update users set money=? where name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")
func main() {
	dsnes := [...]string{
		"admin:lwm_1478963@tcp(193.8.83.217:3306)/rrzcms?charset=utf8mb4&parseTime=True&loc=Local",
	}
	countDb := len(dsnes)
	dbs = make([]*gorm.DB, countDb)
	for i, dsn := range dsnes {
		var err error
		dbs[i], err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("mysql connect error: %s\n", err)
		}
		initDb(dbs[i])
	}

	flag.StringVar(&path, "path", "D:/tmp_CsM7AM/tmp/default", ``)
	flag.IntVar(&sleep, "sleep", 3, "inter time")
	flag.IntVar(&limit, "limit", 3, "limit article pub num")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	successNum := 0

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

			// 这里成功一条就换下一个数据库插入
			dbindex := successNum % countDb
			db := dbs[dbindex]

			if !fi.IsDir() {
				continue
			}

			if successNum >= limit {
				break label1
			}
			if !fi.IsDir() {
				fmt.Printf(" %s 不是有效目录文件 ...... \n", fi.Name())
				continue
			}
			res := pubArticle(fi.Name(), db)
			// 发布成功
			if res {
				// 删除文件
				deleteFile(fi.Name())
				log.Printf("%s publish success \n\n", fi.Name())
				successNum++
			}
			// 是否需要休息几秒
			if sleep > 0 {
				time.Sleep(time.Second * time.Duration(sleep))
			}
		}
		log.Printf("发布完成 数量:%d ...... \n", successNum)
	}()

	// 监听程序退出信号
	s := <-c
	fmt.Println("Got signal:", s)
}

func deleteFile(fileName string) {
	fileName = strings.TrimRight(fileName, "/") + "/"
	if err := os.RemoveAll(path + "/" + fileName); err != nil {
		log.Printf("del dir %s error: %s \n", fileName, err)
	} else {
		log.Printf("delete %s success \n", fileName)
	}
}

// initDb 检查数据库是否初始化
func initDb(db *gorm.DB) {
	var node string
	db.Raw("select name from rrz_article_nodes where id = 25").Scan(&node)
	if node != "外卖系统" {
		fmt.Println("init db ", db.Name())
		db.Exec(`delete from rrz_site_menus where id > 0;`)
		db.Exec(`delete from rrz_article_nodes where id > 24;`)
		// TODO: 这里增删分类后 发布时取随机数也要改
		db.Exec(`INSERT INTO rrz_article_nodes VALUES ('25', '0', '外卖系统', '', 'wmxt', '0', '', '25', '1', '0', '1001', '', '', '{typedir}/{aid}.html', '{typedir}/list_{tid}_{page}.html', '外卖系统', '', '', '&nbsp;', '', 'true', '1639041413');`)
		db.Exec(`INSERT INTO rrz_article_nodes VALUES ('26', '0', '跑腿系统', '', 'ptxt', '0', '', '26', '1', '0', '1002', '', '', '{typedir}/{aid}.html', '{typedir}/list_{tid}_{page}.html', '跑腿系统', '', '', '&nbsp;', '', 'true', '1639041413');`)
		db.Exec(`INSERT INTO rrz_article_nodes VALUES ('27', '0', '微信外卖', '', 'wxwm', '0', '', '27', '1', '0', '1003', '', '', '{typedir}/{aid}.html', '{typedir}/list_{tid}_{page}.html', '微信外卖', '', '', '&nbsp;', '', 'true', '1639041413');`)
		db.Exec(`INSERT INTO rrz_article_nodes VALUES ('28', '0', '同城外卖', '', 'tcwm', '0', '', '28', '1', '0', '1004', '', '', '{typedir}/{aid}.html', '{typedir}/list_{tid}_{page}.html', '同城外卖', '', '', '&nbsp;', '', 'true', '1639041413');`)
		db.Exec(`INSERT INTO rrz_site_menus VALUES ('1', '首页', 'Home', 'index', '0', '1', '1', '/', '1', '1001', 'false', '');`)
		db.Exec(`INSERT INTO rrz_site_menus VALUES ('2', '外卖系统', '', 'wmxt', '0', '2', '1', '/node/25.html', '2', '1002', 'false', '');`)
		db.Exec(`INSERT INTO rrz_site_menus VALUES ('3', '跑腿系统', '', 'ptxt', '0', '3', '1', '/node/26.html', '3', '1003', 'false', '');`)
		db.Exec(`INSERT INTO rrz_site_menus VALUES ('4', '微信外卖', '', 'wxwm', '0', '4', '1', '/node/27.html', '4', '1004', 'false', '');`)
		db.Exec(`INSERT INTO rrz_site_menus VALUES ('5', '同城外卖', '', 'tcwm', '0', '5', '1', '/node/28.html', '5', '1005', 'false', '');`)
	}
}

// pubArticle 发布文章
func pubArticle(articleName string, db *gorm.DB) bool {
	var aid int
	db.Raw("select id from rrz_articles where title = ?", articleName).Scan(&aid)
	if aid > 0 {
		log.Printf("%s 已成功插入(%d)\n", articleName, aid)
		// 删除文件
		deleteFile(articleName)
		return false
	}

	articlePath := path + "/" + articleName

	rand.Seed(time.Now().UnixNano())
	article := Article{
		Title:    articleName,
		NodeID:   rand.Intn(4) + 25, // 随机取
		KeyWords: readFile(articlePath + "/keywords.txt"),
		Content:  dealContent(readFile(articlePath + "/" + articleName + ".txt")),
		Desc:     readFile(articlePath + "/description.txt"),
		// Tags:     readFile(articlePath + "/tags.txt"),
		// Type:     "",
	}

	if article.Content == "" {
		log.Printf("%s 的文章内容为空 ...... \n", articleName)
		return false
	}

	nowInt := time.Now().Unix()
	db.Exec(`insert into rrz_articles (node_id, title, content, uptime, pubtime, ifpub, seo_title, seo_description, seo_keywords, add_time, is_head) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		article.NodeID,
		article.Title,
		article.Content,
		nowInt,
		nowInt,
		"true",
		article.Title,
		article.Desc,
		article.KeyWords,
		nowInt,
		1,
	)
	return true
}

// readFile 读取文件
func readFile(path string) string {
	str, err := ioutil.ReadFile(path)
	if err == nil {
		return string(str)
	}
	return ""
}

// dealContent 文章内容处理 自动分段落
func dealContent(content string) string {
	cs := strings.Split(content, "。")
	if len(cs) > 2 {
		return formatSlice(cs, len(cs)/4, "。")
	}
	cs = strings.Split(content, "，")
	return formatSlice(cs, len(cs)/4, "，")
}

func formatSlice(cs []string, col int, sep string) (str string) {
	var temp string
	for i, v := range cs {
		if len(v) < 2 {
			continue
		}
		if temp != "" && i%col == 0 {
			str += "<p>" + temp + "</p>"
			temp = v + sep
		} else {
			temp += v + sep
		}
	}
	return
}
