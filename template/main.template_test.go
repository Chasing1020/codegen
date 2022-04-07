/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:47
File: main_test.go
*/

package template

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestMainTemplate(t *testing.T) {
	headTemplate, err := template.New("MainTemplate").Parse(MainTemplate)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
