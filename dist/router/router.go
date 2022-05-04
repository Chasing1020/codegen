// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T00:47:23+08:00
// File: response.go

// Package router provide router init function
package router

import (
	"crud/dal"
	"crud/model"
	// _ "crud/docs"
	"crud/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
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

// InitEngine will create gin.Group and add handlers
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
		user.GET("/query", handler.GetUserHandler)
		user.POST("/insert", handler.PostUserHandler)
		user.PUT("/update", handler.PutUserHandler)
		user.DELETE("/delete", handler.DeleteUserHandler)
	}

	book := e.Group("/book")
	{
		book.GET("/query", handler.GetBookHandler)
		book.POST("/insert", handler.PostBookHandler)
		book.PUT("/update", handler.PutBookHandler)
		book.DELETE("/delete", handler.DeleteBookHandler)
	}

	teacher := e.Group("/teacher")
	{
		teacher.GET("/query", handler.GetTeacherHandler)
		teacher.POST("/insert", handler.PostTeacherHandler)
		teacher.PUT("/update", handler.PutTeacherHandler)
		teacher.DELETE("/delete", handler.DeleteTeacherHandler)
	}

	return e
}
