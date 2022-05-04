/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:22
File: router.go
*/

package router

var HeadTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: response.go

// Package router provide router init function
package router

import (
	"{{.Package}}/dal"
	"{{.Package}}/model"
	_ "{{.Package}}/docs"
	"{{.Package}}/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)
`
var MethodsTemplate = `
func init() {
	err := dal.DB.AutoMigrate({{range .Tables}}
		&model.{{.Name}}{},{{end}}
	)
	if err != nil {
		panic(err)
	}
}

// InitEngine will create gin.Group and add handlers
func InitEngine() *gin.Engine {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) { c.String(200, "Hello, world!") })
	e.GET("/ping", func(c *gin.Context) { c.String(200, dal.RDB.Ping(c).Val()) })


	// if you want to use Swagger, please use <swag init> command in the root directory
	// and uncomment the import
	// See "http://localhost:8080/swagger/index.html" for more information
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	{{range .Tables}}
	{{.Tag}} := e.Group("/{{.Tag}}")
	{
		{{.Tag}}.GET("/query", handler.Get{{.Name}}Handler)
		{{.Tag}}.POST("/insert", handler.Post{{.Name}}Handler)
		{{.Tag}}.PUT("/update", handler.Put{{.Name}}Handler)
		{{.Tag}}.DELETE("/delete", handler.Delete{{.Name}}Handler)
	}
	{{end}}
	return e
}
`
