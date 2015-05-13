package web

import (
	"net/http"
	"strings"
)

type handle struct {
	Routers       []*router
	notAllow      handleFunc
	notFound      handleFunc
	staticService handleFunc
}

//函数结构体 无返回值
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
	if isFile {
		//走静态文件服务器
		Logs.Debug("来到了文件服务器")
		h.staticService(c)
		return
	}
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
	return
}

//根据请求找对应的handleFunc
func (h handle) match(r *http.Request) (handleFunc, int, bool) {
	method := r.Method
	path := strings.ToLower(r.URL.Path)
	Logs.Debug("path %s", path)
	for _, router := range h.Routers {
		if router.isFile && strings.HasPrefix(path, router.uri) {
			Logs.Debug("过了match")
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
