
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

// PostBookHandler
// @Description Create Book
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /book/ [post]
// curl --location --request POST 'localhost:8080/book' --form 'Name'='' --form 'Author'='' --form 'Price'='' --form 'Isbn'=''
func PostBookHandler(c *gin.Context) {
	var book model.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&book) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.CreateBook(c, &book)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.CreateBook failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// GetBookHandler
// @Description Query Book
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /book/ [get]
// curl --location --request GET 'localhost:8080/book?ids=1&ids=2&ids=3&limit=3&offset=1'
func GetBookHandler(c *gin.Context) {
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

	books, err := dal.GetBooks(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: books})
}

// PutBookHandler
// @Description Update Book
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /book/ [put]
// curl --location --request PUT 'localhost:8080/book' --form 'id=''' --form 'Name'='' --form 'Author'='' --form 'Price'='' --form 'Isbn'=''
func PutBookHandler(c *gin.Context) {
	var book *model.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&book) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.UpdateBook(c, book)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.UpdateBook(c, book) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: book})
}

// DeleteBookHandler
// @Description Delete Book
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp success
// @Router /book/delete [delete]
// curl --location --request DELETE 'localhost:8080/book?ids=1&ids=2'
func DeleteBookHandler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	books, err := dal.DeleteBooks(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.DeleteBooks(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: books})
}
