package web

import (
	"fmt"
	"net/http"
)

type Controller struct {
	Response http.ResponseWriter
	Request  *http.Request
	//httpNotFound         handleFunc
	//httpMethodNotAllow handleFunc
	Result   map[string]interface{}
	Internal map[string]interface{}
}

func (c *Controller) Write(d string) {
	fmt.Fprintf(c.Response, d)
}

//模板渲染
func (c *Controller) Render(tlpName string) {
	c.render(tlpName)
}

//模板 第一次的时候动态加载
func (c *Controller) render(name string) {
	var err error
	if name[0] != '/' {
		name = "/" + name
	}
	t, ok := sltemplate.Templates[name]
	if !ok || Debug {
		t, err = AddTemplate(name)
		if err != nil {
			Logs.Error("AddTemplate filename:%s err:%v", name, err)
			NotFound(c)
			return
		}
	}
	Logs.Debug("SlTemplates %v", sltemplate.Templates)
	err = t.Execute(c.Response, c.Result)
	if err != nil {
		Logs.Error("Execute err %v", err)
	}

}
