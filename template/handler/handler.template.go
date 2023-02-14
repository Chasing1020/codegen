/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:16
File: handler.go
*/

package handler

var HeadTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: handler.go

// Package handler registers handler functions
package handler

import (
	"{{.Package}}/dal"
	"{{.Package}}/model"
	"github.com/gin-gonic/gin"
)
`

var MethodsTemplate = `
// Insert{{.Name}}Handler godoc
// @Summary      Insert {{.Name}}
// @Description  Insert {{.Name}}
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        {{.Name}}  body    model.{{.Name}} true "{{.Name}}"
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /{{.SnakeCase}}/insert [post]
func Insert{{.Name}}Handler(c *gin.Context) {
	var {{.LowerCamelCase}} model.{{.Name}}
	err := c.ShouldBind(&{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.LowerCamelCase}}) failed: " + err.Error()})
		return
	}

	err = dal.Create{{.Name}}(c.Request.Context(), &{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.Create{{.Name}} failed: " + err.Error()})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}})
}

// Delete{{.Name}}Handler godoc
// @Summary      Delete {{.Name}}
// @Description  Delete {{.Name}}
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        {{.Name}}  body    model.{{.Name}} true "{{.Name}}"
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /{{.SnakeCase}}/delete [post]
func Delete{{.Name}}Handler(c *gin.Context) {
	var {{.LowerCamelCase}} *model.{{.Name}}
	err := c.ShouldBind(&{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.LowerCamelCase}}) failed: " + err.Error()})
		return
	}

	{{.LowerCamelCase}}s, err := dal.Delete{{.Name}}s(c.Request.Context(), {{.LowerCamelCase}})

	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "func dal.Delete{{.Name}}s(c.Request.Context(), ids) failed: " + err.Error()})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}s})
}

// Update{{.Name}}Handler godoc
// @Summary      Update {{.Name}}
// @Description  Update {{.Name}}:
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        {{.LowerCamelCase}}  body    model.{{.Name}} true "{{.LowerCamelCase}}"
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Router       /{{.SnakeCase}}/update [post]
func Update{{.Name}}Handler(c *gin.Context) {
	var {{.LowerCamelCase}} *model.{{.Name}}
	err := c.ShouldBind(&{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.LowerCamelCase}}) failed: " + err.Error()})
		return
	}

	err = dal.Update{{.Name}}(c.Request.Context(), {{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.Update{{.Name}}(c.Request.Context(), {{.LowerCamelCase}}) failed:" + err.Error()})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}})
}

// Get{{.Name}}Handler godoc
// @Summary      Get {{.Name}}
// @Description  Get {{.Name}}:
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        id        query   int         true   "id"
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Failure      404       {object}  model.Resp  "failed"
// @Router       /{{.SnakeCase}}/get [get]
func Get{{.Name}}Handler(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(400, model.Resp{Code: 400, Message: "c.GetQuery(id) failed"})
		return
	}

	{{.LowerCamelCase}}s, err := dal.Get{{.Name}}ById(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, model.Resp{Code: 404, Message: "not found: " + err.Error()})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}s})
}

// Query{{.Name}}Handler godoc
// @Summary      Query {{.Name}}
// @Description  Query {{.Name}}:
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        {{.LowerCamelCase}}  body    model.{{.Name}}Param true "{{.LowerCamelCase}}Param"
// @Success      200       {object}  model.Resp  "success"
// @Failure      400       {object}  model.Resp  "failed"
// @Failure      404       {object}  model.Resp  "failed"
// @Router       /{{.SnakeCase}}/query [post]
func Query{{.Name}}Handler(c *gin.Context) {
	var req *model.{{.Name}}Param
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.LowerCamelCase}}) failed: " + err.Error()})
		return
	}

	{{.LowerCamelCase}}s, count, err := dal.Query{{.Name}}s(c.Request.Context(), req.{{.Name}}, req.Limit, req.Offset, req.NeedCount)
	if err != nil {
		c.JSON(404, model.Resp{Code: 404, Message: "not found: " + err.Error()})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}s, Count: count})
}
`
