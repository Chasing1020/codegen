/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:59
File: case.go
*/

package config

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(camel string) (snake string) {
	snake = matchFirstCap.ReplaceAllString(camel, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake =  strings.ToLower(snake)
	return
}

func ToUpperCamelCase(snake string) (camel string) {
	isToUpper := false
	for k, v := range snake {
		if k == 0 {
			camel = strings.ToUpper(string(snake[0]))
		} else {
			if isToUpper {
				camel += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camel += string(v)
				}
			}
		}
	}
	return
}

func ToLowerCamelCase(snake string) (camel string) {
	camel = ToUpperCamelCase(snake)
	camel = strings.ToLower(camel[:1]) + camel[1:]
	return
}