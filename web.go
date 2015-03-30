package web

import (
	"net/http"
	"strings"
	"time"
	"github.com/astaxie/beego/logs"
    "fmt"
)

var (
	Addr           string
	Routers        []*router
	NotAllow	handleFunc
	NotFound	handleFunc
	Logs		*logs.BeeLogger
	Debug		bool
    TemplatePath string
)

type App struct {
	handle *handle
	addr    string
	debug	bool

}

type handle struct {
	Routers []*router
	notAllow handleFunc
	notFound handleFunc
}

// return true is return
type handleFunc func(*Controller)

//control router
func (h handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Controller{
		Response: w,
		Request:  r,
		Result:   make(map[string]interface{}),
		Internal: make(map[string]interface{}),
	}
	handle, status, isFile := h.match(r)
	if status == http.StatusOK {
		handle(c)
		Logs.Info("%s %s", r.Method, r.URL)
		return
	}
	if status == http.StatusNotFound {
		h.notFound(c)
		Logs.Info("%s %s Not Found", r.Method, r.URL)
		return
	}
	if status == http.StatusMethodNotAllowed {
        Logs.Info("%s %s 404", r.Method, r.URL)
        h.notAllow(c)
        return
    }
	if isFile {
		return
	}
	return
}

func (h handle) match(r *http.Request) (handleFunc, int, bool) {
	method := r.Method
	path := strings.ToLower(r.URL.Path)
	for _, router := range(h.Routers) {
		if router.isFile && strings.HasPrefix(path, router.uri) {
			return router.handle, http.StatusOK, true
		}
		if path == router.uri {
            if router.method == "" || in(method, strings.ToUpper(router.method)) {
				return router.handle, http.StatusOK, false
			} else {
				return router.handle, http.StatusMethodNotAllowed, false
			}
		}
	}
	return nil, http.StatusNotFound, false
}

func Router(uri string, method string, handleFunc handleFunc) {
	thisRouter := &router{
		uri:    uri,
		method: method,
		isFile: false,
		handle: handleFunc,
	}
	Routers = append(Routers, thisRouter)
}

func Init() *App {
	defaultHandler := &handle{
		Routers: Routers,
		notAllow: NotAllow,
		notFound: NotFound,
	}
	app := &App{
		handle: defaultHandler,
		addr:    Addr,
		debug:	Debug,
	}
	return app
}

func init(){
    NotAllow = defaultNotAllow
    NotFound = defaultNotFound
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

func defaultNotFound(c *Controller) {
    fmt.Fprintf(c.Response, "404")
    return
}

func defaultNotAllow(c *Controller) {
    fmt.Fprintf(c.Response, "mothod notAllow")
    return
}

//TODO 静态页面
//TODO 静态模板
