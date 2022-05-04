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

// PostBookHandler godoc
// @Summary      Insert Book
// @Description  Insert Book:
// @Description  curl --location --request POST 'localhost:8080/book/insert' --form 'name'='' --form 'author'='' --form 'price'='' --form 'isbn'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        book  body    model.Book true "Book"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /book/insert [post]
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

// GetBookHandler godoc
// @Summary      Query Book
// @Description  Query Book:
// @Description  curl --location --request GET 'localhost:8080/book/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /book/query [get]
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

// PutBookHandler godoc
// @Summary      Update Book
// @Description  Update Book:
// @Description  curl --location --request PUT 'localhost:8080/book' --form 'id'='' --form 'name'='' --form 'author'='' --form 'price'='' --form 'isbn'=''
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        book  body    model.Book true "Book"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /book/update [put]
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

// DeleteBookHandler godoc
// @Summary      Delete Book
// @Description  Delete Book
// @Description  curl --location --request DELETE 'localhost:8080/book/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids       query   int         true   "id list"
// @Success      200       object  model.Resp  success
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /book/delete [delete]
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
