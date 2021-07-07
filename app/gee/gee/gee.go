package gee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// 实现通用的Logger中间件，能够记录请求到响应所花费的时间
type Engine struct {
	*RouterGroup
	groups []*RouterGroup // map无序会导致错误
	*router
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

type RouterGroup struct {
	prefix      string       // 路由前缀
	middlewares []HandFunc   // 中间件
	parent      *RouterGroup // 链表 父节点
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	routerGroup := &RouterGroup{engine: engine}
	engine.RouterGroup = routerGroup
	engine.groups = append(engine.groups, routerGroup)
	return engine
}

func (rg *RouterGroup) Group(name string) *RouterGroup {
	nrg := &RouterGroup{
		prefix: rg.prefix + name,
		parent: rg,
		engine: rg.engine,
	}
	nrg.engine.groups = append(nrg.engine.groups, nrg)
	return nrg
}

func (rg *RouterGroup) addRoute(method string, pattern string, handler HandFunc) {
	// 真实路由 = 前缀 + 路由
	pattern = rg.prefix + pattern
	rg.engine.router.addRoute(method, pattern, handler)
}

func (rg *RouterGroup) Use(handler ...HandFunc) *RouterGroup {
	rg.middlewares = append(rg.middlewares, handler...)
	return rg
}

// GET defines the method to add GET request
func (rg *RouterGroup) GET(pattern string, handler HandFunc) {
	rg.addRoute("GET", pattern, handler)
}

func (rg *RouterGroup) POST(pattern string, handler HandFunc) {
	rg.addRoute("POST", pattern, handler)
}

func (g *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := newContext(w, r)
	context.index = -1
	context.engine = g
	g.passMiddleware(context)
	g.router.handle(context)
}

func (g *Engine) passMiddleware(c *Context) {
	for _, rg := range g.groups {
		if strings.HasPrefix(c.Req.URL.Path, rg.prefix) {
			if rg.middlewares != nil {
				c.handlers = append(c.handlers, rg.middlewares...)
			}
		}
	}
}

func (g *Engine) Run(address string) {
	if err := http.ListenAndServe(address, g); err != nil {
		log.Fatalf("listen on %s error: %s \n", address, err)
	}
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}
