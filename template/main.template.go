/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:42
File: main.go
*/

package template

var MainTemplate = `/*
Copyright © 2022 {{.Author}} <{{.Email}}>
Time: {{.Time}}
File: response.go
*/

package main

import (
	"{{.Package}}/router"
)

// @title {{.Package}}
// @version 1.0
// @description {{.Package}} Project
// @contact.name {{.Author}}
// @contact.url https://github.com/{{.Author}}
// @contact.email {{.Email}}
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	engine := router.InitEngine()
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}`

