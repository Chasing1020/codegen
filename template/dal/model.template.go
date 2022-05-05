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
// Get{{.Name}}s will query {{.Name}} by ids, limit and offset
func Get{{.Name}}s(ctx context.Context, ids []string, limit int, offset int) ([]model.{{.Name}}, error) {
	var {{.Tag}}s []model.{{.Name}}
	if len(ids) == 0 {
		return nil, nil
	}
	conn := DB.WithContext(ctx)

	if limit <= 0 && limit >= 500 {
		limit = 100
	}
	conn = conn.Limit(limit)
	if offset > 0 {
		conn = conn.Offset(offset)
	}

	err := conn.Find(&{{.Tag}}s, ids).Error
	if err != nil {
		log.Println("func Get{{.Name}}s failed: ", err)
		return nil, err
	}
	return {{.Tag}}s, nil
}

// Create{{.Name}} will create a(n) {{.Name}} by *model.{{.Name}}
func Create{{.Name}}(ctx context.Context, {{.Tag}} *model.{{.Name}}) error {
	conn := DB.WithContext(ctx)
	err := conn.Create(&{{.Name}}).Error
	if err != nil {
		log.Println("func Create{{.Name}} failed: ", err)
		return err
	}
	return nil
}

// Update{{.Name}} will update a(n) {{.Name}} by *model.{{.Name}}.ID and set the value to *model.{{.Name}}
func Update{{.Name}}(ctx context.Context, {{.Tag}} *model.{{.Name}}) error {
	conn := DB.WithContext(ctx)
	err := conn.Where("id = ?", {{.Name}}.ID).Updates({{.Name}}).Error
	if err != nil {
		log.Println("func Update{{.Name}} failed: ", err)
		return err
	}
	return nil
}

// Delete{{.Name}}s will delete all {{.Name}} by ids
func Delete{{.Name}}s(ctx context.Context, ids []string) ([]model.{{.Name}}, error) {
	var {{.Tag}}s []model.{{.Name}}
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&{{.Tag}}s).Error
	if err != nil {
		log.Println("func Delete{{.Name}}s failed: ", err)
		return nil, err
	}
	return {{.Tag}}s, nil
}`
