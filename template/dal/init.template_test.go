/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-16:43
File: init_test.go
*/

package dal

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestInit(t *testing.T) {
	tpl, err := template.New("InitTemplate").Parse(InitTemplate)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
