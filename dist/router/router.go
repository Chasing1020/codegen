/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-04-07 23:11:10.330178 +0800 CST m=+0.001950585
File: response.go
*/

package router

import (
	"crud/dal"
	"crud/model"
	// _ "crud/docs"
	"crud/handler"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := dal.DB.AutoMigrate(
		&model.User{},
		&model.Book{},
		&model.Teacher{},
	)
	if err != nil {
		panic(err)
	}
}

func InitEngine() *gin.Engine {
	e := gin.Default()
	e.GET("/", func(c *gin.Context) { c.String(200, "Hello, world!") })
	e.GET("/ping", func(c *gin.Context) { c.String(200, dal.RDB.Ping(c).Val()) })
	
	// if you want to use Swagger, please use <swag init> command in the root directory
	// and uncomment the import
	// See "http://localhost:8080/swagger/index.html" for more information
	// e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	
	user := e.Group("/user")
	{
		user.GET("/", handler.GetUserHandler)
		user.POST("/", handler.PostUserHandler)
		user.PUT("/", handler.PutUserHandler)
		user.DELETE("/", handler.DeleteUserHandler)
	}
	
	book := e.Group("/book")
	{
		book.GET("/", handler.GetBookHandler)
		book.POST("/", handler.PostBookHandler)
		book.PUT("/", handler.PutBookHandler)
		book.DELETE("/", handler.DeleteBookHandler)
	}
	
	teacher := e.Group("/teacher")
	{
		teacher.GET("/", handler.GetTeacherHandler)
		teacher.POST("/", handler.PostTeacherHandler)
		teacher.PUT("/", handler.PutTeacherHandler)
		teacher.DELETE("/", handler.DeleteTeacherHandler)
	}
	
	return e
}
