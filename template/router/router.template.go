/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:22
File: router.go
*/

package router

var HeadTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: router.go

// Package router provide router init function
package router

import (
	"{{.Package}}/dal"
	"{{.Package}}/model"
	"{{.Package}}/auth"
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
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(CORS())
	// If you want to use Swagger, please use <swag fmt && swag init> command in the root directory
	// See "http://localhost:8080/swagger/index.html" for more information
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	actuator := e.Group("/actuator")
	{
		actuator.GET("/health", HealthHandler)
		actuator.GET("/health/redis", RedisHealthHandler)
		actuator.GET("/health/mysql", MySQLHealthHandler)
		actuator.Use(auth.Session).GET("/health/session", SessionHealthHandler)
	}

	e.Use(gin.Logger())
	// e.Use(auth.Session) // enable session authentication
	
	// e.POST("/login", auth.Login)
	// e.GET("/logout", auth.Logout)

	{{range .Tables}}
	{{.LowerCamelCase}} := e.Group("/{{.SnakeCase}}")
    // {{.LowerCamelCase}}.Use(auth.CookieRequired) // uncomment to require authentication to access
	{
		{{.LowerCamelCase}}.GET("/get", handler.Get{{.Name}}Handler)
		{{.LowerCamelCase}}.POST("/query", handler.Query{{.Name}}Handler)
		{{.LowerCamelCase}}.POST("/insert", handler.Insert{{.Name}}Handler)
		{{.LowerCamelCase}}.POST("/update", handler.Update{{.Name}}Handler)
		{{.LowerCamelCase}}.POST("/delete", handler.Delete{{.Name}}Handler)
	}
	{{end}}
	return e
}



func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("Origin"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
`