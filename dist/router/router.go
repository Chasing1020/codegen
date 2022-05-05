// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: router.go

// Package router provide router init function
package router

import (
	"crud/auth"
	"crud/dal"
	_ "crud/docs"
	"crud/handler"
	"crud/model"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	err := dal.DB.AutoMigrate(
		&model.Course{},
		&model.CourseOutline{},
		&model.Department{},
		&model.Elective{},
		&model.Student{},
	)
	if err != nil {
		panic(err)
	}
}

// InitEngine will create gin.Group and add handlers
func InitEngine() *gin.Engine {
	e := gin.Default()

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

	// e.Use(auth.Session) // enable session authentication
	// e.POST("/login", auth.Login)
	// e.GET("/logout", auth.Logout)

	course := e.Group("/course")
	// course.Use(auth.CookieRequired) // uncomment to require authentication to access
	{
		course.GET("/query", handler.GetCourseHandler)
		course.POST("/insert", handler.PostCourseHandler)
		course.PUT("/update", handler.PutCourseHandler)
		course.DELETE("/delete", handler.DeleteCourseHandler)
	}

	courseOutline := e.Group("/course_outline")
	// courseOutline.Use(auth.CookieRequired) // uncomment to require authentication to access
	{
		courseOutline.GET("/query", handler.GetCourseOutlineHandler)
		courseOutline.POST("/insert", handler.PostCourseOutlineHandler)
		courseOutline.PUT("/update", handler.PutCourseOutlineHandler)
		courseOutline.DELETE("/delete", handler.DeleteCourseOutlineHandler)
	}

	department := e.Group("/department")
	// department.Use(auth.CookieRequired) // uncomment to require authentication to access
	{
		department.GET("/query", handler.GetDepartmentHandler)
		department.POST("/insert", handler.PostDepartmentHandler)
		department.PUT("/update", handler.PutDepartmentHandler)
		department.DELETE("/delete", handler.DeleteDepartmentHandler)
	}

	elective := e.Group("/elective")
	// elective.Use(auth.CookieRequired) // uncomment to require authentication to access
	{
		elective.GET("/query", handler.GetElectiveHandler)
		elective.POST("/insert", handler.PostElectiveHandler)
		elective.PUT("/update", handler.PutElectiveHandler)
		elective.DELETE("/delete", handler.DeleteElectiveHandler)
	}

	student := e.Group("/student")
	// student.Use(auth.CookieRequired) // uncomment to require authentication to access
	{
		student.GET("/query", handler.GetStudentHandler)
		student.POST("/insert", handler.PostStudentHandler)
		student.PUT("/update", handler.PutStudentHandler)
		student.DELETE("/delete", handler.DeleteStudentHandler)
	}

	return e
}
