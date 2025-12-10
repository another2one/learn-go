package main

// 创建 http 服务器

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	_ "io/ioutil"
	"learn-go/app/http_hello_world/utils"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed template/*
var staticFS embed.FS
var tmplFiles = template.Must(template.ParseFS(staticFS, "template/*.html"))

// MyHandler 路由处理
type MyHandler struct{}

var sid string

//var body []byte

func goBadu(w http.ResponseWriter) {
	// 读取百度页面
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic("read from baidu error")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

// SendTemplate 发送模板
func SendTemplate(w http.ResponseWriter) {
	value := utils.Session.Get(sid, "name").(string)
	value += "-1"
	utils.Session.Set(sid, "name", value)
	err := tmplFiles.ExecuteTemplate(w, "about.html", value)
	if err != nil {
		fmt.Println("execute template error: ", err)
	}
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 响应头部
	w.Header().Add("lizhi", "666")

	utils.Logger.Info("请求方法：", r.Method)
	utils.Logger.Errorf("err:%s", r.URL.Path)

	// cookie
	if cookie, err := r.Cookie("sid"); errors.Is(err, http.ErrNoCookie) {
		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "sid",
			Value:    utils.Session.GetSessionID(),
			MaxAge:   utils.Session.ExpireSeconds,
		}
		//w.Header().Set("Set-Cookie", cookie.String())
		sid = cookie.Value
		http.SetCookie(w, &cookie)
	} else {
		sid = cookie.Value
	}

	// 重定向
	//w.Header().Add("Location", "https://www.baidu.com")
	//w.WriteHeader(http.StatusFound)

	// 响应体
	SendTemplate(w)

	// 请求体
	requestBody := make([]byte, r.ContentLength)
	r.Body.Read(requestBody)
}

// streamHandler 提供基本的流式文本数据
func streamHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头，允许跨域和流式内容
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// 检查响应是否支持Flusher接口
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// 启用分块传输
	w.Header().Set("Transfer-Encoding", "chunked")

	// 模拟持续数据流
	for i := 1; i <= 5; i++ {
		// 构建消息
		message := fmt.Sprintf("Data chunk #%d at %s\n", i, time.Now().Format("15:04:05"))
		fmt.Fprintf(w, message)
		flusher.Flush()             // 立即将数据块发送到客户端
		time.Sleep(1 * time.Second) // 每秒发送一次
	}

	// 发送结束信号
	fmt.Fprintf(w, "Stream completed at %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// Message 定义流式传输的数据结构
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Type    string `json:"type"`  // 数据类型，如 'info', 'warning', 'error'
	Value   int    `json:"value"` // 用于图表显示的数值
}

// metricsStreamHandler 提供JSON格式的指标数据流，用于前端图表
func metricsStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// 模拟实时系统指标（CPU、内存、网络等）
	types := []string{"cpu", "memory", "network"}

	for i := 1; i <= 50; i++ {
		// 为每种类型生成一个数据点
		for _, metricType := range types {
			msg := Message{
				ID:      i,
				Content: fmt.Sprintf("%s usage update", metricType),
				Time:    time.Now().Format("15:04:05"),
				Type:    metricType,
				Value:   generateRandomValue(metricType),
			}

			jsonData, err := json.Marshal(msg)
			if err != nil {
				log.Printf("JSON marshaling error: %v", err)
				continue
			}

			fmt.Fprintf(w, "%s\n", jsonData)
			flusher.Flush()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// generateRandomValue 根据指标类型生成模拟数值
func generateRandomValue(metricType string) int {
	// 基于类型生成合理的随机值
	switch metricType {
	case "cpu":
		return 20 + rand.IntN(80) // CPU使用率 20%-99%
	case "memory":
		return 30 + rand.IntN(80) // 内存使用率 30%-99%
	case "network":
		return 50 + rand.IntN(80) // 网络流量
	default:
		return rand.IntN(100)
	}
}

func main() {

	server := http.NewServeMux()

	bootstrap()

	// 注册处理请求的两种方式
	// 第一种 HandleFunc
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	goBadu(w)
	//})

	// 第二种 Handle
	server.Handle("/666", &MyHandler{})

	// 处理静态资源
	server.Handle("/template/static/", http.FileServer(http.FS(staticFS)))

	// 响应流
	server.HandleFunc("/chunk", streamHandler)
	server.HandleFunc("/sse", metricsStreamHandler)

	server.HandleFunc("/query", func(writer http.ResponseWriter, request *http.Request) {
		utils.Query()
		_, err := writer.Write([]byte("query"))
		if err != nil {
			fmt.Println("write error:", err)
		}
	})

	// 创建监听，也可用 server 方法进行更加详细的配置
	log.Printf("About to listen on 8080. Go to http://127.0.0.1:8080/")
	http1 := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	go func() {
		err := http1.ListenAndServe()
		if err != nil {
			panic("start http err:" + err.Error())
		}
	}()

	listenSignal(context.Background(), http1)
}

func listenSignal(ctx context.Context, httpSrv *http.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		close()
		log.Println("http shutdown")
	}
}

func bootstrap() {
	utils.InitLogger()
	utils.InitDB()
	utils.InitSession()
}

func close() {
	log.Println("close ...")
	utils.Logger.Sync()
	utils.Session.Close()
}
