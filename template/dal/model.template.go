/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:45
File: model.go
*/

package dal

var HeadTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: model.go

// Package dal represents as Data access layer
package dal

import (
	"{{.Package}}/model"
	"context"
	"log"
)
`


var MethodsTemplate = `
// Create{{.Name}} will create a(n) {{.Name}} by *model.{{.Name}}
func Create{{.Name}}(ctx context.Context, param *model.{{.Name}}) error {
	err := DB.WithContext(ctx).Create(&param).Error
	if err != nil {
		log.Println("func Create{{.Name}} failed: ", err)
		return err
	}
	return nil
}

// Delete{{.Name}}s will delete {{.Name}} by param
func Delete{{.Name}}s(ctx context.Context, param *model.{{.Name}}) ([]*model.{{.Name}}, error) {
	if param == nil {return nil, nil}
	var {{.LowerCamelCase}}s []*model.{{.Name}}

	session := DB.WithContext(ctx)
	if param.ID != 0 {
		session = session.Where("id = ?", param.ID)
	}{{range .Columns}}
	if param.{{.Name}} != {{.DefaultValue}} {
		session = session.Where("{{.SnakeCase}} = ?", param.{{.Name}})
	}{{end}}

	err := session.Delete(&{{.LowerCamelCase}}s).Error
	if err != nil {
		log.Println("func Get{{.Name}}s failed: ", err)
		return nil, err
	}
	return {{.LowerCamelCase}}s, nil
}

// Update{{.Name}} will update a(n) {{.Name}} by *model.{{.Name}}.ID and set the value to *model.{{.Name}}
func Update{{.Name}}(ctx context.Context, param *model.{{.Name}}) error {
	err := DB.WithContext(ctx).Where("id = ?", param.ID).Updates(param).Error
	if err != nil {
		log.Println("func Update{{.Name}} failed: ", err)
		return err
	}
	return nil
}

// Get{{.Name}}ById will query {{.Name}} by id
func Get{{.Name}}ById(ctx context.Context, id string) (*model.{{.Name}}, error) {
	var {{.LowerCamelCase}} *model.{{.Name}}
	if id == "0" { return nil, nil }
	err := DB.WithContext(ctx).First(&{{.LowerCamelCase}}, id).Error
	if err != nil {
		log.Println("func Get{{.Name}}ById failed: ", err)
		return nil, err
	}
	return {{.LowerCamelCase}}, nil
}

// Query{{.Name}}s will query {{.Name}} by given Parameters
func Query{{.Name}}s(ctx context.Context, param *model.{{.Name}}, limit, offset int, needCount bool) ([]*model.{{.Name}}, int64, error) {
	if param == nil { return nil, 0, nil }
	var {{.LowerCamelCase}}s []*model.{{.Name}}
	var count int64

	session := DB.WithContext(ctx)
	if param.ID != 0 {
		session = session.Where("id = ?", param.ID)
	}{{range .Columns}}
	if param.{{.Name}} != {{.DefaultValue}} {
		session = session.Where("{{.SnakeCase}} = ?", param.{{.Name}})
	}{{end}}

	if needCount {
		err := session.Model(&param).Count(&count).Error
		if err != nil {
			return nil, 0, err
		}
	}

	if limit <= 0 || limit >= 500 {
		limit = 100
	}
	session = session.Limit(limit)
	if offset > 0 {
		session = session.Offset(offset)
	}
	result := session.Find(&{{.LowerCamelCase}}s)
	if result.Error != nil {
		log.Println("func Get{{.Name}}s failed: ", result.Error)
		return nil, 0, result.Error
	}
	return {{.LowerCamelCase}}s, count, nil
}
`
