package web

import (
	"errors"
	"html/template"
)

//模板  子模板  模板函数

var (
	sltemplate *SlTemplate
)

type SlTemplate struct {
	BaseTemplate []string //储存 路径[文件内容]
	Templates    map[string]*template.Template
	FuncMap      template.FuncMap
}

func NewSlTemplate() *SlTemplate {
	var tmpList []string
	return &SlTemplate{
		BaseTemplate: tmpList,
		Templates:    make(map[string]*template.Template),
		FuncMap:      make(template.FuncMap),
	}
}

func init() {
	sltemplate = NewSlTemplate()
}

//添加 模板函数
func AddFuncMap(key string, funcName interface{}) error {
	if _, ok := sltemplate.FuncMap[key]; ok {
		return errors.New("this template func already have!")
	} else {
		sltemplate.FuncMap[key] = funcName
	}
	return nil
}

//添加基础模板

func AddBaseTemplate(htmlPath string) error {
	if htmlPath == "" {
		return errors.New("htmlpath err")
	}
	sltemplate.BaseTemplate = append(sltemplate.BaseTemplate, htmlPath)
	return nil
}

//添加模板
func AddTemplate(name string) (t *template.Template, err error) {

	//基本数据
	t, err = template.ParseFiles(TemplatePath + name)
	Logs.Debug(TemplatePath + name)
	if err != nil {
		Logs.Error("template file %s not found err:%v", name, err)
	}
	//子模板
	for _, v := range sltemplate.BaseTemplate {
		t, err = t.ParseFiles(TemplatePath + v)
		if err != nil {
			Logs.Error("template file %s not found err:%v", v, err)
		}
	}
	//模板函数
	t = t.Funcs(sltemplate.FuncMap)
	sltemplate.Templates[name] = t
	return
}
