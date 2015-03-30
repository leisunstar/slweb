package web

import (
    "html/template"
)

var (
    slTemplateFuncMap template.FuncMap
    SlTemplates map[string]*template.Template

)

func init(){
    SlTemplates = make(map[string]*template.Template)
    slTemplateFuncMap = make(template.FuncMap)
}

func AddFuncMap(key string, funname interface{}) error {
    slTemplateFuncMap[key] = funname
    return nil
}

func AddTemplate(name string)(t *template.Template, err error){


    t, err = template.ParseFiles(TemplatePath+name)
    Logs.Debug(TemplatePath+name)
    if err != nil{
        Logs.Error("template file %s not found err:%v", name, err)
    }
    SlTemplates[name] = t
    return
}