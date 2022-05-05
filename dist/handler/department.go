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

// PostDepartmentHandler godoc
// @Summary      Insert Department
// @Description  Insert Department:
// @Description  curl --location --request POST 'localhost:8080/department/insert' --form 'departmentId'='' --form 'name'='' --form 'address'='' --form 'phone'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Department  body    model.Department  true  "Department"
// @Success      200         object  model.Resp        success
// @Failure      400         object  model.Resp        failed
// @Router       /department/insert [post]
func PostDepartmentHandler(c *gin.Context) {
	var department model.Department
	err := c.ShouldBind(&department)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&department) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateDepartment(c, &department)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateDepartment failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetDepartmentHandler godoc
// @Summary      Query Department
// @Description  Query Department:
// @Description  curl --location --request GET 'localhost:8080/department/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200     object  model.Resp  success
// @Failure      400     object  model.Resp  failed
// @Router       /department/query [get]
func GetDepartmentHandler(c *gin.Context) {
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

	departments, err := dal.GetDepartments(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: departments})
}

// PutDepartmentHandler godoc
// @Summary      Update Department
// @Description  Update Department:
// @Description  curl --location --request PUT 'localhost:8080/department/update' --form 'id'='' --form 'departmentId'='' --form 'name'='' --form 'address'='' --form 'phone'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Department  body    model.Department  true  "Department"
// @Success      200         object  model.Resp        success
// @Failure      400         object  model.Resp        failed
// @Router       /department/update [put]
func PutDepartmentHandler(c *gin.Context) {
	var department *model.Department
	err := c.ShouldBind(&department)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&department) failed:" + err.Error(), Data: nil})
		return
	}
	if department.ID == 0 {
		c.JSON(400, model.Resp{Code: 400, Message: "parameter 'id' required", Data: nil})
		return
	}

	err = dal.UpdateDepartment(c, department)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateDepartment(c, department) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: department})
}

// DeleteDepartmentHandler godoc
// @Summary      Delete Department
// @Description  Delete Department
// @Description  curl --location --request DELETE 'localhost:8080/department/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids  query   int         true  "id list"
// @Success      200  object  model.Resp  success
// @Success      200  object  model.Resp  success
// @Failure      400  object  model.Resp  failed
// @Router       /department/delete [delete]
func DeleteDepartmentHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	departments, err := dal.DeleteDepartments(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteDepartments(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: departments})
}
