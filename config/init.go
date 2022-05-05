/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-22:20
File: init.go
*/

package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"strings"
	"time"
)

var Conf Configuration
var wd string

func init() {
	data, err := os.ReadFile(ProjectPath() + "/config/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)

	Conf.Module.Time = time.Now()
	for _, table := range Conf.Schema.Tables {
		table.Name = ToUpperCamelCase(table.Name)
		table.Tag = strings.ToLower(table.Name[:1]) + table.Name[1:]
		for _, column := range table.Columns {
			column.Name = ToUpperCamelCase(column.Name)
			column.Tag = ToSnakeCase(column.Name)
		}
	}
}

func ProjectPath() string {
	if wd != "" {
		return wd
	}
	var err error
	wd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(wd + "/config"); os.IsNotExist(err) {
			wd = path.Join(wd, "/..")
		} else {
			break
		}
	}
	return wd
}
