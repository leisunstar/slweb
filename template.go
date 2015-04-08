package web

import (
    "html/template"
)

var (
    sltemplate *SlTemplate
)

type SlTemplate struct {
    BaseTemplate   map[string]string
    Templates      map[string]*template.Template
    FuncMap        template.FuncMap
}

func NewSlTemplate()(*SlTemplate){
    return &SlTemplate{
        BaseTemplate: make(map[string]string, 0),
        Templates: make(map[string]*template.Template),
        FuncMap : make(template.FuncMap),
    }
}

func init(){
    sltemplate = NewSlTemplate()
}

func AddFuncMap(key string, funcName interface{}) error {
    sltemplate.FuncMap[key] = funcName
    return nil
}

func AddTemplate(name string)(t *template.Template, err error){


    t, err = template.ParseFiles(TemplatePath+name)
    Logs.Debug(TemplatePath+name)
    if err != nil{
        Logs.Error("template file %s not found err:%v", name, err)
    }
    sltemplate.Templates[name] = t
    return
}