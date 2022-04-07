/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-19:38
File: common.go
*/

package gen

import (
	"codegen/config"
	"os"
	"text/template"
)

func WriteHeadCode(fileName, headTemplate string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	ExecuteTemplate(file, headTemplate, config.Conf.Module)

	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func WriteMethodsCode(fileName, headTemplate, methodsTemplate string, table *config.Table) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	ExecuteTemplate(file, headTemplate, config.Conf.Module)
	ExecuteTemplate(file, methodsTemplate, table)

	err = file.Sync()
	if err != nil {
		panic(err)
	}
}


func WriteRouterCode(fileName, headTemplate, methodsTemplate string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	ExecuteTemplate(file, headTemplate, config.Conf.Module)
	ExecuteTemplate(file, methodsTemplate, config.Conf)

	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func ExecuteTemplate(file *os.File, templateText string, data interface{}) {
	methods, err := template.New(file.Name()).Parse(templateText)
	if err != nil {
		panic(err)
	}
	err = methods.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
