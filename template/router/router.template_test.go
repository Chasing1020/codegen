/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:28
File: router_test.go
*/

package router

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestRouter(t *testing.T) {
	headTemplate, err := template.New("HeadTemplate").Parse(HeadTemplate)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}

	methodsTemplate, err := template.New("MethodsTemplate").Parse(MethodsTemplate)
	if err != nil {
		panic(err)
	}
	err = methodsTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
