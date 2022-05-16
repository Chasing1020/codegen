// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T19:02:15+08:00
// File: handler.go

// Package handler registers handler functions
package handler

import (
	"crud/dal"
	"crud/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostStudentHandler godoc
// @Summary      Insert Student
// @Description  Insert Student:
// @Description  curl --location --request POST 'localhost:8080/student/insert' --form 'studentId'='' --form 'password'='' --form 'name'='' --form 'sex'='' --form 'birthday'='' --form 'hometown'='' --form 'phone'='' --form 'departmentId'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Student  body    model.Student  true  "Student"
// @Success      200      object  model.Resp     success
// @Failure      400      object  model.Resp     failed
// @Router       /student/insert [post]
func PostStudentHandler(c *gin.Context) {
	var student model.Student
	err := c.ShouldBind(&student)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&student) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateStudent(c, &student)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateStudent failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetStudentHandler godoc
// @Summary      Query Student
// @Description  Query Student:
// @Description  curl --location --request GET 'localhost:8080/student/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200     object  model.Resp  success
// @Failure      400     object  model.Resp  failed
// @Router       /student/query [get]
func GetStudentHandler(c *gin.Context) {
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

	students, err := dal.GetStudents(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: students})
}

// PutStudentHandler godoc
// @Summary      Update Student
// @Description  Update Student:
// @Description  curl --location --request PUT 'localhost:8080/student/update' --form 'id'='' --form 'studentId'='' --form 'password'='' --form 'name'='' --form 'sex'='' --form 'birthday'='' --form 'hometown'='' --form 'phone'='' --form 'departmentId'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Student  body    model.Student  true  "Student"
// @Success      200      object  model.Resp     success
// @Failure      400      object  model.Resp     failed
// @Router       /student/update [put]
func PutStudentHandler(c *gin.Context) {
	var student *model.Student
	err := c.ShouldBind(&student)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&student) failed:" + err.Error(), Data: nil})
		return
	}
	if student.ID == 0 {
		c.JSON(400, model.Resp{Code: 400, Message: "parameter 'id' required", Data: nil})
		return
	}

	err = dal.UpdateStudent(c, student)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateStudent(c, student) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: student})
}

// DeleteStudentHandler godoc
// @Summary      Delete Student
// @Description  Delete Student
// @Description  curl --location --request DELETE 'localhost:8080/student/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids  query   int         true  "id list"
// @Success      200  object  model.Resp  success
// @Success      200  object  model.Resp  success
// @Failure      400  object  model.Resp  failed
// @Router       /student/delete [delete]
func DeleteStudentHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	students, err := dal.DeleteStudents(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteStudents(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: students})
}
