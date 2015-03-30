package web

import ()


type router struct {
    uri    string
    method string
    isFile bool
    handle handleFunc
}
