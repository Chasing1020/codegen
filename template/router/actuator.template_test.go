/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/5-15:59
File: actuator_test.go
*/

package router

import (
	"codegen/config"
	"os"
	"testing"
	"text/template"
)



func TestActuator(t *testing.T) {
	headTemplate, err := template.New("HeadTemplate").Parse(ActuatorHeadTemplate)
	if err != nil {
		panic(err)
	}
	err = headTemplate.Execute(os.Stdout, config.Conf.Module)
	if err != nil {
		panic(err)
	}
	//
	//methodsTemplate, err := template.New("MethodsTemplate").Parse(MethodsTemplate)
	//if err != nil {
	//	panic(err)
	//}
	//err = methodsTemplate.Execute(os.Stdout, config.Conf.Module)
	//if err != nil {
	//	panic(err)
	//}
}

