package web

import ()

type router struct {
	uri    string     //地址
	method string     //Post Get
	isFile bool       //是不是文件
	handle handleFunc //controller 函数
}

//添加router
func Router(uri string, method string, handleFunc handleFunc) {
	thisRouter := &router{
		uri:    uri,
		method: method,
		isFile: false,
		handle: handleFunc,
	}
	Routers = append(Routers, thisRouter)
}
