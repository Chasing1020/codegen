/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:25
File: response_test.go
*/

package model

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)

func TestResponse(t *testing.T) {
	headTemplate, err := template.New("HeadTemplate").Parse(ResponseTemplate)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
}
