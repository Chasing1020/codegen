/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-16:39
File: parse_test.go
*/

package conf

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestParse(t *testing.T) {
	tpl, err := template.New("YamlTemplate").Parse(YamlTemplate)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
