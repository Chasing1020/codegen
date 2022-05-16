// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022/5/16T11:38:19+08:00
// File: model_mock.template.go

package dal

var MockHeadTemplate = `
// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: model.go

package dal

import (
	"context"
	"{{.Package}}/model"
	"{{.Package}}/dal"
	"testing"
)

`
var MockMethodTemplate = `
func TestMock{{.Name}}(t *testing.T) {
	ctx := context.Background()

	{{.LowerCamelCase}}s := make([]*model.{{.Name}}, 20)
	for _, {{.LowerCamelCase}} := range {{.LowerCamelCase}}s {
		{{.LowerCamelCase}} = &model.{{.Name}}{
			// TODO: fill all fields
		}
		dal.Create{{.Name}}(ctx, {{.LowerCamelCase}})
	}
}
`
