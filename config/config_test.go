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
	"strings"
	"testing"
	"time"
)

func TestToLower(t *testing.T) {
	fmt.Println(ToSnakeCase("Name"))
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)
	Conf.Module.Time = time.Now()
	for _, table := range Conf.Schema.Tables {
		table.Name = strings.Title(table.Name)
		table.Tag = ToSnakeCase(table.Name)
		for _, column := range table.Columns {
			column.Name = strings.Title(column.Name)
			column.Tag = ToSnakeCase(column.Name)
		}
	}
	fmt.Println(Conf)

}


func TestGetPath(t *testing.T) {
	fmt.Println(ProjectPath())
}