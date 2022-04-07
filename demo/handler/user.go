/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:42
File: user.go
*/

package handler

import (
	"codegen/dal"
	"codegen/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostUserHandler
// @Description Create User
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /user/ [post]
// curl --location --request POST 'localhost:8080/user' \
// --form 'name=""' \
// --form 'password=""'
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

// GetUserHandler
// @Description Query User
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /user/ [get]
// curl --location --request GET 'localhost:8080/user?ids=1&ids=2'
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

// PutUserHandler
// @Description Update User
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /user/ [put]
// curl --location --request PUT 'localhost:8080/user' \
// --form 'id=""' \
// --form 'name=""' \
// --form 'password=""'
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

// DeleteUserHandler
// @Description Delete User
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp success
// @Router /user/delete [delete]
// curl --location --request DELETE 'localhost:8080/user?ids=1&ids=2'
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
