/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-19:45
File: generator.go
*/

package gen

import (
	"codegen/config"
	"codegen/template"
	"codegen/template/conf"
	"codegen/template/dal"
	"codegen/template/handler"
	"codegen/template/model"
	"codegen/template/router"
)

var path = config.ProjectPath()

var Functions = []func(){
	Conf, Dal, Handler, Model, Router, Main,
}

func Conf() {
	WriteHeadCode(path+"/dist/conf/config.go", conf.Template)
	WriteHeadCode(path+"/dist/conf/config.yaml", conf.YamlTemplate)
}

func Dal() {
	WriteHeadCode(path+"/dist/dal/init.go", dal.InitTemplate)
	for _, table := range config.Conf.Tables {
		WriteMethodsCode(path+"/dist/dal/"+table.Tag+".go",
			dal.HeadTemplate, dal.MethodsTemplate, table)
	}
}

func Handler() {
	for _, table := range config.Conf.Tables {
		WriteMethodsCode(path+"/dist/handler/"+table.Tag+".go",
			handler.HeadTemplate, handler.MethodsTemplate, table)
	}
}

func Model() {
	WriteHeadCode(path+"/dist/model/response.go", model.ResponseTemplate)
	for _, table := range config.Conf.Tables {
		WriteMethodsCode(path+"/dist/model/"+table.Tag+".go",
			model.HeadTemplate, model.StructTemplate, table)
	}
}

func Router() {
	WriteRouterCode(path+"/dist/router/router.go", router.HeadTemplate, router.MethodsTemplate)
}

func Main() {
	WriteHeadCode(path+"/dist/main.go", template.MainTemplate)
	WriteHeadCode(path+"/dist/go.mod", template.GoModTemplate)
}
