package main

// 创建 http 服务器

import (
	"app/http_hello_world/utils"
	"fmt"
	"html/template"
	_ "io/ioutil"
	"log"
	"net/http"
)

type MyHandler struct{}

//var body []byte

//func init (){
//	// 读取百度页面
//	resp, err := http.Get("http://www.baidu.com")
//	if err != nil {
//		panic("read from baidu error")
//		return
//	}
//	defer resp.Body.Close()
//	body, _ = ioutil.ReadAll(resp.Body)
//	fmt.Println(string(body))
//}

func SendTemplate(w http.ResponseWriter) {
	t := template.Must(template.ParseGlob("../template/*.html"))
	//t.Execute(w, "666") // 解析第一个模板
	t.ExecuteTemplate(w, "about.html", "666") // 解析指定的模板文件
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 响应头部
	w.Header().Add("lizhi", "666")

	// cookie
	if _, err := r.Cookie("sid"); err == http.ErrNoCookie {
		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "sid",
			Value:    utils.Session.GetSessionId(),
			MaxAge:   utils.Session.ExpireTime,
		}
		//w.Header().Set("Set-Cookie", cookie.String())
		http.SetCookie(w, &cookie)
	}

	// 重定向
	//w.Header().Add("Location", "https://www.baidu.com")
	//w.WriteHeader(http.StatusFound)

	// 响应体
	//fmt.Fprint(w, "666")
	//w.Write(body)
	SendTemplate(w)

	// 请求行
	url := r.URL
	fmt.Println("url: ", url.Scheme, url.Host, url.Path, url.RawQuery, url.String(), url.Query())

	// 请求头
	requestHeader := r.Header
	fmt.Printf("requestHeader: %+v \n", requestHeader["Accept"])

	// 请求体
	requestBody := make([]byte, r.ContentLength)
	r.Body.Read(requestBody)
	fmt.Printf("requestBody: %v \n", string(requestBody))
	// 请求参数
	r.ParseForm() // FormValue() 和 PostFormValue() 会隐式调用该方法
	// 所有：如果请求头与post里面含有相同键，则值会添加到一个数组里面去，而且表单数据在前面
	fmt.Printf("params: %v \n", r.Form)
	// post application/x-www-form-urlencoded 才支持调用， application/form-data 需要用 MultipartForm
	fmt.Printf("form: %v \n", r.PostForm)

}

func main() {

	// 注册处理请求的两种方式
	// 第一种 HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ...
	})

	// 第二种 Handle
	http.Handle("/666", &MyHandler{})

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080/")

	// 处理静态资源
	// 去掉 /static/ 前缀后转到 ../template/static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../template/static"))))

	// 创建监听，也可用 server 方法进行更加详细的配置
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("创建服务器错误：", err)
	}

}
