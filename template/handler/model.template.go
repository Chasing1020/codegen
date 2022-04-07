/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:16
File: model.go
*/

package handler

var HeadTemplate = `
/*
Copyright © 2022 {{.Author}} <{{.Email}}>
Time: {{.Time}}
File: handler.go
*/

package handler

import (
	"{{.Package}}/dal"
	"{{.Package}}/model"
	"github.com/gin-gonic/gin"
	"strconv"
)
`

var MethodsTemplate = `
// Post{{.Name}}Handler
// @Description Create {{.Name}}
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /{{.Tag}}/ [post]
// curl --location --request POST 'localhost:8080/{{.Tag}}'{{range .Columns}} --form '{{.Name}}'=''{{end}}
func Post{{.Name}}Handler(c *gin.Context) {
	var {{.Tag}} model.{{.Name}}
	err := c.ShouldBind(&{{.Tag}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.Tag}}) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.Create{{.Name}}(c, &{{.Tag}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "del.Create{{.Name}} failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: nil})
}

// Get{{.Name}}Handler
// @Description Query {{.Name}}
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /{{.Tag}}/ [get]
// curl --location --request GET 'localhost:8080/{{.Tag}}?ids=1&ids=2&ids=3&limit=3&offset=1'
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

	{{.Tag}}s, err := dal.Get{{.Name}}s(c, ids, limit, offset)

	if err != nil {
		c.JSON(404, model.Resp{Code: 400, Message: "not found: " + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.Tag}}s})
}

// Put{{.Name}}Handler
// @Description Update {{.Name}}
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp failed
// @Router /{{.Tag}}/ [put]
// curl --location --request PUT 'localhost:8080/{{.Tag}}' --form 'id='''{{range .Columns}} --form '{{.Name}}'=''{{end}}
func Put{{.Name}}Handler(c *gin.Context) {
	var {{.Tag}} *model.{{.Name}}
	err := c.ShouldBind(&{{.Tag}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "c.ShouldBind(&{{.Tag}}) failed:" + err.Error(), Data: nil})
		return
	}

	err = dal.Update{{.Name}}(c, {{.Tag}})
	if err != nil {
		c.JSON(400, model.Resp{Code: 400, Message: "dal.Update{{.Name}}(c, {{.Tag}}) failed:" + err.Error(), Data: nil})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.Tag}}})
}

// Delete{{.Name}}Handler
// @Description Delete {{.Name}}
// @Accept application/json
// @Accept application/x-www-form-urlencoded
// @Success 200 object model.Resp success
// @Failure 400 object model.Resp success
// @Router /{{.Tag}}/delete [delete]
// curl --location --request DELETE 'localhost:8080/{{.Tag}}?ids=1&ids=2'
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

	{{.Tag}}s, err := dal.Delete{{.Name}}s(c, ids)

	if err != nil {
		c.JSON(400, model.Resp{
			Code:    400,
			Message: "func dal.Delete{{.Name}}s(c, ids) failed:" + err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(200, model.Resp{Code: 200, Message: "success", Data: {{.Tag}}s})
}
`