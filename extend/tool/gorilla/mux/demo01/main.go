package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// golang自带路由库 http.ServerMux ，实际上是一个 map[string]Handler，是请求的url路径和该url路径对于的一个处理函数的映射关系.缺点：
// 1. 不支持参数设定，例如/user/:uid 这种泛型类型匹配
// 2. 无法很友好的支持REST模式，无法限制访问方法（POST，GET等）
// 3. 也不支持正则
func main() {
	r := mux.NewRouter()

	// 1.普通路由
	r.HandleFunc("/", IndexHandler)

	// 2. 正则路由参数，下面例子中限制为英文字母
	r.HandleFunc("/articles/{title:[a-z]+}", TitleHandler)

	// 3. 路由匹配
	// 	3.1 host匹配
	//只匹配 www.example.com/sub
	h := r.Host("www.example.com").Subrouter()
	h.HandleFunc("/sub", IndexHandler)
	// 动态匹配子路由
	//  *.example.com/ac
	h1 := r.Host("{subdomain:[a-z]+}.example.com").Subrouter()
	h1.HandleFunc("/ac", IndexHandler)

	// 3.2 更多的一些其他匹配
	r.PathPrefix("/products/").
		Methods("GET", "POST").                                      //请求方法匹配
		Schemes("https").                                            //schemes
		Headers("X-Requested-With", "XMLHttpRequest").               //header 匹配
		Queries("key", "value").                                     //query的值匹配
		MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool { // 用户自定义方法 匹配
			return r.ProtoMajor == 0
		})

	// 3.4 匹配的顺序 添加顺序匹配

	// 4. 分组 Subrouter 实现

	// 5. 中间件
	h.Use(loggingMiddleware)

	// 6. 遍历注册的所有路由
	h.Walk(walk)

	// 7. 命名路由
	h.HandleFunc("/sub/{name:[a-z|1-9]+}", IndexHandler).Name("subname")
	url1, err := r.Get("subname").URL()
	if err != nil {
		log.Println("get url1 error:", err)
	} else {
		log.Printf("url1 = %v \n", url1)
	}
	url2, err := r.Get("subname").URL("name", "lizhi")
	if err != nil {
		log.Println("get url2 error:", err)
	} else {
		log.Println("url2 = ", url2)
	}

	log.Println("visit: localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln("listen err:", err)
	}
}

// localhost:8080
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}

// localhost:8080/articles/title/lizhi
//  只能英文，其他404
func TitleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // 获取参数
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "title: %v\n", vars["title"])
}

// Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Do stuff here
		fmt.Println(r.RequestURI)
		fmt.Fprintf(w, "%s\r\n", r.URL)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func walk(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	pathTemplate, err := route.GetPathTemplate()
	if err == nil {
		fmt.Println("ROUTE:", pathTemplate)
	}
	pathRegexp, err := route.GetPathRegexp()
	if err == nil {
		fmt.Println("Path regexp:", pathRegexp)
	}
	queriesTemplates, err := route.GetQueriesTemplates()
	if err == nil {
		fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
	}
	queriesRegexps, err := route.GetQueriesRegexp()
	if err == nil {
		fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
	}
	methods, err := route.GetMethods()
	if err == nil {
		fmt.Println("Methods:", strings.Join(methods, ","))
	}
	fmt.Println()
	return nil
}
