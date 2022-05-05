/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-19:45
File: generator.go
*/

package gen

import (
	"codegen/config"
	"codegen/template"
	"codegen/template/auth"
	"codegen/template/conf"
	"codegen/template/dal"
	"codegen/template/handler"
	"codegen/template/model"
	"codegen/template/router"
	"log"
)

var path = config.ProjectPath()

var Functions = []func(){
	Auth, Conf, Dal, Handler, Model, Router, Main,
}

func Auth() {
	WriteHeadCode(path+"/dist/auth/session.go", auth.Template)
	log.Println(path + "/dist/auth/ generated")
}

func Conf() {
	WriteHeadCode(path+"/dist/conf/config.go", conf.Template)
	WriteHeadCode(path+"/dist/conf/config.yaml", conf.YamlTemplate)
	log.Println(path + "/dist/conf/ generated")
}

func Dal() {
	WriteHeadCode(path+"/dist/dal/init.go", dal.InitTemplate)
	for _, table := range config.Conf.Tables {
		WriteMethodsCode(path+"/dist/dal/"+config.ToSnakeCase(table.SnakeCase)+".go",
			dal.HeadTemplate, dal.MethodsTemplate, table)
	}
	log.Println(path + "/dist/dal/ generated")
}

func Handler() {
	for _, table := range config.Conf.Tables {
		WriteMethodsCode(path+"/dist/handler/"+config.ToSnakeCase(table.SnakeCase)+".go",
			handler.HeadTemplate, handler.MethodsTemplate, table)
	}
	log.Println(path + "/dist/handler/ generated")
}

func Model() {
	WriteHeadCode(path+"/dist/model/response.go", model.ResponseTemplate)
	for _, table := range config.Conf.Tables {
		WriteMethodsCode(path+"/dist/model/"+config.ToSnakeCase(table.SnakeCase)+".go",
			model.HeadTemplate, model.StructTemplate, table)
	}
	log.Println(path + "/dist/model/ generated")
}

func Router() {
	WriteHeadCode(path+"/dist/router/actuator.go", router.ActuatorHeadTemplate)
	WriteRouterCode(path+"/dist/router/router.go", router.HeadTemplate, router.MethodsTemplate)
	log.Println(path + "/dist/router/ generated")
}

func Main() {
	WriteHeadCode(path+"/dist/main.go", template.MainTemplate)
	WriteHeadCode(path+"/dist/go.mod", template.GoModTemplate)
	WriteHeadCode(path+"/dist/Makefile", template.MakefileTemplate)
	log.Println(path + "/dist/ generated")
}
