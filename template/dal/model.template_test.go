/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-16:49
File: model_test.go
*/

package dal

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestModel(t *testing.T) {

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
	err = methodsTemplate.Execute(os.Stdout, config.Conf.Tables[0])
	if err != nil {
		panic(err)
	}
}
