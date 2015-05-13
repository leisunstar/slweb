package web

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"time"
)

var (
	Addr          string //地址
	Routers       []*router
	NotAllow      handleFunc
	NotFound      handleFunc
	StaticService handleFunc
	Logs          *logs.BeeLogger
	Debug         bool
	TemplatePath  string
	StaticPath    string //绝对路径
)

type App struct {
	handle *handle
	addr   string
	debug  bool
}

//添加静态文件目录
func Static(staticPath, staticfilePath string) {
	StaticPath = staticfilePath
	r := &router{
		uri:    staticPath,
		isFile: true,
	}
	Routers = append(Routers, r)
}

func Init() *App {
	defaultHandler := &handle{
		Routers:       Routers,
		notAllow:      NotAllow,
		notFound:      NotFound,
		staticService: StaticService,
	}
	app := &App{
		handle: defaultHandler,
		addr:   Addr,
		debug:  Debug,
	}
	return app
}

//默认404
func defaultNotFound(c *Controller) {
	fmt.Fprintf(c.Response, "404")
	return
}

//默认无此方法
func defaultNotAllow(c *Controller) {
	fmt.Fprintf(c.Response, "mothod notAllow")
	return
}

//默认静态处理
func defaultStatic(c *Controller) {
	handler := http.FileServer(http.Dir(StaticPath))
	handler.ServeHTTP(c.Response, c.Request)
	return
}

func init() {
	//先将404 等赋给默认方法
	NotAllow = defaultNotAllow
	NotFound = defaultNotFound
	StaticService = defaultStatic
}

// run server in addr
func (a *App) Run() error {
	server := &http.Server{
		Addr:           a.addr,
		Handler:        a.handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	Logs.Info("Run Server %s", a.addr)
	return server.ListenAndServe()
}

