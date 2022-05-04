/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-05-05T00:47:23+08:00
File: handler.go
*/

// Package handler registers handler functions
package handler

import (
	"crud/dal"
	"crud/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostTeacherHandler godoc
// @Summary      Insert Teacher
// @Description  Insert Teacher:
// @Description  curl --location --request POST 'localhost:8080/teacher/insert' --form 'name'='' --form 'department'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        teacher  body    model.Teacher true "Teacher"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /teacher/insert [post]
func PostTeacherHandler(c *gin.Context) {
	var teacher model.Teacher
	err := c.ShouldBind(&teacher)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&teacher) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateTeacher(c, &teacher)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateTeacher failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetTeacherHandler godoc
// @Summary      Query Teacher
// @Description  Query Teacher:
// @Description  curl --location --request GET 'localhost:8080/teacher/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /teacher/query [get]
func GetTeacherHandler(c *gin.Context) {
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

	teachers, err := dal.GetTeachers(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: teachers})
}

// PutTeacherHandler godoc
// @Summary      Update Teacher
// @Description  Update Teacher:
// @Description  curl --location --request PUT 'localhost:8080/teacher' --form 'id'='' --form 'name'='' --form 'department'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        teacher  body    model.Teacher true "Teacher"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /teacher/update [put]
func PutTeacherHandler(c *gin.Context) {
	var teacher *model.Teacher
	err := c.ShouldBind(&teacher)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&teacher) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.UpdateTeacher(c, teacher)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateTeacher(c, teacher) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: teacher})
}

// DeleteTeacherHandler godoc
// @Summary      Delete Teacher
// @Description  Delete Teacher
// @Description  curl --location --request DELETE 'localhost:8080/teacher/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids       query   int         true   "id list"
// @Success      200       object  model.Resp  success
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /teacher/delete [delete]
func DeleteTeacherHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	teachers, err := dal.DeleteTeachers(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteTeachers(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: teachers})
}
