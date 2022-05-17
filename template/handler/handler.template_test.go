/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:18
File: model_test.go
*/

package handler

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestHandler(t *testing.T) {

	headTemplate, err := template.New("HeadTemplate").Parse(HeadTemplate)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}

	methodsTemplate, err := template.New("ModelHandlerMethodsTemplate").Parse(MethodsTemplate)
	if err != nil {
		panic(err)
	}
	err = methodsTemplate.Execute(os.Stdout, config.Conf.Tables[0])
	if err != nil {
		panic(err)
	}
}
