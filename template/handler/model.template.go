/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:16
File: model.go
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
	"strconv"
)
`

var MethodsTemplate = `
// Post{{.Name}}Handler godoc
// @Summary      Insert {{.Name}}
// @Description  Insert {{.Name}}:
// @Description  curl --location --request POST 'localhost:8080/{{.SnakeCase}}/insert'{{range .Columns}} --form '{{.LowerCamelCase}}'=''{{end}}
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        {{.Name}}  body    model.{{.Name}} true "{{.Name}}"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /{{.SnakeCase}}/insert [post]
func Post{{.Name}}Handler(c *gin.Context) {
	var {{.LowerCamelCase}} model.{{.Name}}
	err := c.ShouldBind(&{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.LowerCamelCase}}) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.Create{{.Name}}(c, &{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.Create{{.Name}} failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// Get{{.Name}}Handler godoc
// @Summary      Query {{.Name}}
// @Description  Query {{.Name}}:
// @Description  curl --location --request GET 'localhost:8080/{{.SnakeCase}}/query?ids=1&ids=2&ids=3&limit=3&offset=1'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids     query   int         true   "id list"
// @Param        limit   query   int         false  "limit"   default(0)
// @Param        offset  query   int         false  "offset"  default(0)
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /{{.SnakeCase}}/query [get]
func Get{{.Name}}Handler(c *gin.Context) {
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

	{{.LowerCamelCase}}s, err := dal.Get{{.Name}}s(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}s})
}

// Put{{.Name}}Handler godoc
// @Summary      Update {{.Name}}
// @Description  Update {{.Name}}:
// @Description  curl --location --request PUT 'localhost:8080/{{.SnakeCase}}/update' --form 'id'=''{{range .Columns}} --form '{{.LowerCamelCase}}'=''{{end}}
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        {{.Name}}  body    model.{{.Name}} true "{{.Name}}"
// @Success      200       object  model.Resp  success
// @Failure      400       object  model.Resp  failed
// @Router       /{{.SnakeCase}}/update [put]
func Put{{.Name}}Handler(c *gin.Context) {
	var {{.LowerCamelCase}} *model.{{.Name}}
	err := c.ShouldBind(&{{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.LowerCamelCase}}) failed:" + err.Error(), Data: nil})
		return
	}
	if {{.LowerCamelCase}}.ID == 0 {
		c.JSON(400, model.Resp{Code: 400, Message: "parameter 'id' required", Data: nil})
		return
	}

	err = dal.Update{{.Name}}(c, {{.LowerCamelCase}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.Update{{.Name}}(c, {{.LowerCamelCase}}) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}})
}

// Delete{{.Name}}Handler godoc
// @Summary      Delete {{.Name}}
// @Description  Delete {{.Name}}
// @Description  curl --location --request DELETE 'localhost:8080/{{.SnakeCase}}/delete?ids=1&ids=2'
// @Accept       application/json
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Param        ids       query   int         true   "id list"
// @Success      200       object  model.Resp  success
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /{{.SnakeCase}}/delete [delete]
func Delete{{.Name}}Handler(c *gin.Context) {
	ids, ok := c.GetQueryArray("ids")
	if !ok {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "c.GetQueryArray(\"ids\") failed.",
			Data:    nil,
		})
		return
	}

	{{.LowerCamelCase}}s, err := dal.Delete{{.Name}}s(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.Delete{{.Name}}s(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.LowerCamelCase}}s})
}
`