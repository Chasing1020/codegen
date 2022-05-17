/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:52
File: init_test.go
*/

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
	"time"
)

func TestToLower(t *testing.T) {
	data, err := os.ReadFile(ProjectPath() + "/config/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)

	Conf.Module.Time = time.Now()
	for _, table := range Conf.Schema.Tables {
		table.Name = ToUpperCamelCase(table.Name)
		table.LowerCamelCase =  ToLowerCamelCase(table.Name)
		table.SnakeCase = ToSnakeCase(table.Name)
		for _, column := range table.Columns {
			column.Name = ToUpperCamelCase(column.Name)
			column.LowerCamelCase = ToLowerCamelCase(column.Name)
			column.SnakeCase = ToSnakeCase(column.Name)
			column.DefaultValue = ToDefaultValue(column.Type)
		}
	}
	fmt.Println(Conf)

}


func TestGetPath(t *testing.T) {
	fmt.Println(ProjectPath())
}