/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:07
File: conf_test.go
*/

package conf

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestConf(t *testing.T) {
	headTemplate, err := template.New("Template").Parse(Template)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
