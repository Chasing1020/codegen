/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:15
File: model_test.go
*/

package model

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

	for _, table := range config.Conf.Schema.Tables {
		methodsTemplate, err := template.New("StructTemplate").Parse(StructTemplate)
		if err != nil {
			panic(err)
		}
		err = methodsTemplate.Execute(os.Stdout, table)
		if err != nil {
			panic(err)
		}
	}
}
