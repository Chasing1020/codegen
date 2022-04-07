
/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-04-07 23:11:10.330178 +0800 CST m=+0.001950585
File: handler.go
*/

package handler

import (
	"crud/dal"
	"crud/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostTeacherHandler
// @Description Create Teacher
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /teacher/ [post]
// curl --location --request POST 'localhost:8080/teacher' --form 'Name'='' --form 'Department'=''
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

// GetTeacherHandler
// @Description Query Teacher
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /teacher/ [get]
// curl --location --request GET 'localhost:8080/teacher?ids=1&ids=2&ids=3&limit=3&offset=1'
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

// PutTeacherHandler
// @Description Update Teacher
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /teacher/ [put]
// curl --location --request PUT 'localhost:8080/teacher' --form 'id=''' --form 'Name'='' --form 'Department'=''
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

// DeleteTeacherHandler
// @Description Delete Teacher
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp success
// @Router /teacher/delete [delete]
// curl --location --request DELETE 'localhost:8080/teacher?ids=1&ids=2'
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
