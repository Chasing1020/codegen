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

// PostCourseOutlineHandler godoc
// @Summary      Insert CourseOutline
// @Description  Insert CourseOutline:
// @Description  curl --location --request POST 'localhost:8080/course_outline/insert' --form 'term'='' --form 'courseId'='' --form 'teacherId'='' --form 'classSchedule'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        CourseOutline  body    model.CourseOutline  true  "CourseOutline"
// @Success      200            object  model.Resp           success
// @Failure      400            object  model.Resp           failed
// @Router       /course_outline/insert [post]
func PostCourseOutlineHandler(c *gin.Context) {
	var courseOutline model.CourseOutline
	err := c.ShouldBind(&courseOutline)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&courseOutline) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateCourseOutline(c, &courseOutline)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateCourseOutline failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetCourseOutlineHandler godoc
// @Summary      Query CourseOutline
// @Description  Query CourseOutline:
// @Description  curl --location --request GET 'localhost:8080/course_outline/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200     object  model.Resp  success
// @Failure      400     object  model.Resp  failed
// @Router       /course_outline/query [get]
func GetCourseOutlineHandler(c *gin.Context) {
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

	courseOutlines, err := dal.GetCourseOutlines(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: courseOutlines})
}

// PutCourseOutlineHandler godoc
// @Summary      Update CourseOutline
// @Description  Update CourseOutline:
// @Description  curl --location --request PUT 'localhost:8080/course_outline/update' --form 'id'='' --form 'term'='' --form 'courseId'='' --form 'teacherId'='' --form 'classSchedule'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        CourseOutline  body    model.CourseOutline  true  "CourseOutline"
// @Success      200            object  model.Resp           success
// @Failure      400            object  model.Resp           failed
// @Router       /course_outline/update [put]
func PutCourseOutlineHandler(c *gin.Context) {
	var courseOutline *model.CourseOutline
	err := c.ShouldBind(&courseOutline)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&courseOutline) failed:" + err.Error(), Data: nil})
		return
	}
	if courseOutline.ID == 0 {
		c.JSON(400, model.Resp{Code: 400, Message: "parameter 'id' required", Data: nil})
		return
	}

	err = dal.UpdateCourseOutline(c, courseOutline)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateCourseOutline(c, courseOutline) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: courseOutline})
}

// DeleteCourseOutlineHandler godoc
// @Summary      Delete CourseOutline
// @Description  Delete CourseOutline
// @Description  curl --location --request DELETE 'localhost:8080/course_outline/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids  query   int         true  "id list"
// @Success      200  object  model.Resp  success
// @Success      200  object  model.Resp  success
// @Failure      400  object  model.Resp  failed
// @Router       /course_outline/delete [delete]
func DeleteCourseOutlineHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	courseOutlines, err := dal.DeleteCourseOutlines(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteCourseOutlines(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: courseOutlines})
}
