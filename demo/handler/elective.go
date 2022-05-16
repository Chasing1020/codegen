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

// PostElectiveHandler godoc
// @Summary      Insert Elective
// @Description  Insert Elective:
// @Description  curl --location --request POST 'localhost:8080/elective/insert' --form 'studentId'='' --form 'term'='' --form 'courseId'='' --form 'teacherId'='' --form 'score'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Elective  body    model.Elective  true  "Elective"
// @Success      200       object  model.Resp      success
// @Failure      400       object  model.Resp      failed
// @Router       /elective/insert [post]
func PostElectiveHandler(c *gin.Context) {
	var elective model.Elective
	err := c.ShouldBind(&elective)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&elective) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateElective(c, &elective)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateElective failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetElectiveHandler godoc
// @Summary      Query Elective
// @Description  Query Elective:
// @Description  curl --location --request GET 'localhost:8080/elective/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200     object  model.Resp  success
// @Failure      400     object  model.Resp  failed
// @Router       /elective/query [get]
func GetElectiveHandler(c *gin.Context) {
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

	electives, err := dal.GetElectives(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: electives})
}

// PutElectiveHandler godoc
// @Summary      Update Elective
// @Description  Update Elective:
// @Description  curl --location --request PUT 'localhost:8080/elective/update' --form 'id'='' --form 'studentId'='' --form 'term'='' --form 'courseId'='' --form 'teacherId'='' --form 'score'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        Elective  body    model.Elective  true  "Elective"
// @Success      200       object  model.Resp      success
// @Failure      400       object  model.Resp      failed
// @Router       /elective/update [put]
func PutElectiveHandler(c *gin.Context) {
	var elective *model.Elective
	err := c.ShouldBind(&elective)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&elective) failed:" + err.Error(), Data: nil})
		return
	}
	if elective.ID == 0 {
		c.JSON(400, model.Resp{Code: 400, Message: "parameter 'id' required", Data: nil})
		return
	}

	err = dal.UpdateElective(c, elective)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateElective(c, elective) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: elective})
}

// DeleteElectiveHandler godoc
// @Summary      Delete Elective
// @Description  Delete Elective
// @Description  curl --location --request DELETE 'localhost:8080/elective/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids  query   int         true  "id list"
// @Success      200  object  model.Resp  success
// @Success      200  object  model.Resp  success
// @Failure      400  object  model.Resp  failed
// @Router       /elective/delete [delete]
func DeleteElectiveHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	electives, err := dal.DeleteElectives(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteElectives(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: electives})
}
