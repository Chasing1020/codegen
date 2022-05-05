// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: handler.go

// Package handler registers handler functions
package handler

import (
	"crud/dal"
	"crud/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostCourseHandler godoc
// @Summary      Insert Course
// @Description  Insert Course:
// @Description  curl --location --request POST 'localhost:8080/course/insert' --form 'courseId'='' --form 'name'='' --form 'credit'='' --form 'time'='' --form 'departmentId'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Course  body    model.Course  true  "Course"
// @Success      200     object  model.Resp    success
// @Failure      400     object  model.Resp    failed
// @Router       /course/insert [post]
func PostCourseHandler(c *gin.Context) {
	var course model.Course
	err := c.ShouldBind(&course)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&course) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateCourse(c, &course)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateCourse failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetCourseHandler godoc
// @Summary      Query Course
// @Description  Query Course:
// @Description  curl --location --request GET 'localhost:8080/course/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200     object  model.Resp  success
// @Failure      400     object  model.Resp  failed
// @Router       /course/query [get]
func GetCourseHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{Code: 400, Message: "c.GetQueryArray(\"ids\") failed", Data: nil})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "cast limit failed: " + err.Error(), Data: nil})
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "cast offset failed: " + err.Error(), Data: nil})
		return
	}

	courses, err := dal.GetCourses(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: courses})
}

// PutCourseHandler godoc
// @Summary      Update Course
// @Description  Update Course:
// @Description  curl --location --request PUT 'localhost:8080/course/update' --form 'id'='' --form 'courseId'='' --form 'name'='' --form 'credit'='' --form 'time'='' --form 'departmentId'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Course  body    model.Course  true  "Course"
// @Success      200     object  model.Resp    success
// @Failure      400     object  model.Resp    failed
// @Router       /course/update [put]
func PutCourseHandler(c *gin.Context) {
	var course *model.Course
	err := c.ShouldBind(&course)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&course) failed:" + err.Error(), Data: nil})
		return
	}
	if course.ID == 0 {
		c.JSON(400, model.Resp{Code: 400, Message: "parameter 'id' required", Data: nil})
		return
	}

	err = dal.UpdateCourse(c, course)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateCourse(c, course) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: course})
}

// DeleteCourseHandler godoc
// @Summary      Delete Course
// @Description  Delete Course
// @Description  curl --location --request DELETE 'localhost:8080/course/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids  query   int         true  "id list"
// @Success      200  object  model.Resp  success
// @Success      200  object  model.Resp  success
// @Failure      400  object  model.Resp  failed
// @Router       /course/delete [delete]
func DeleteCourseHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	courses, err := dal.DeleteCourses(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteCourses(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: courses})
}
