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

// PostUserHandler godoc
// @Summary      Insert User
// @Description  Insert User:
// @Description  curl --location --request POST 'localhost:8080/user/insert' --form 'name'='' --form 'password'='' --form 'grade'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        user  body    model.User true "User"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /user/insert [post]
func PostUserHandler(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&user) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateUser(c, &user)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateUser failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetUserHandler godoc
// @Summary      Query User
// @Description  Query User:
// @Description  curl --location --request GET 'localhost:8080/user/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /user/query [get]
func GetUserHandler(c *gin.Context) {
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

	users, err := dal.GetUsers(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: users})
}

// PutUserHandler godoc
// @Summary      Update User
// @Description  Update User:
// @Description  curl --location --request PUT 'localhost:8080/user' --form 'id'='' --form 'name'='' --form 'password'='' --form 'grade'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        user  body    model.User true "User"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /user/update [put]
func PutUserHandler(c *gin.Context) {
	var user *model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&user) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.UpdateUser(c, user)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateUser(c, user) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: user})
}

// DeleteUserHandler godoc
// @Summary      Delete User
// @Description  Delete User
// @Description  curl --location --request DELETE 'localhost:8080/user/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids       query   int         true   "id list"
// @Success      200       object  model.Resp  success
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /user/delete [delete]
func DeleteUserHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	users, err := dal.DeleteUsers(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteUsers(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: users})
}
