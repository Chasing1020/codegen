/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:48
File: mod_test.go
*/

package template

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestMod(t *testing.T) {
	headTemplate, err := template.New("GoModTemplate").Parse(GoModTemplate)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
