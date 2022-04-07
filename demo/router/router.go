/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:14
File: router.go
*/

package router

import (
	"codegen/dal"
	_ "codegen/docs"
	"codegen/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitEngine() *gin.Engine {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) { c.String(200, "Hello, world!") })
	e.GET("/ping", func(c *gin.Context) { c.String(200, dal.RDB.Ping(c).Val()) })


	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := e.Group("/user")
	{
		user.GET("/", handler.GetUserHandler)
		user.POST("/", handler.PostUserHandler)
		user.PUT("/", handler.PutUserHandler)
		user.DELETE("/", handler.DeleteUserHandler)
	}

	return e
}
