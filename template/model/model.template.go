/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:13
File: model.go
*/

package model

var HeadTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: response.go

// Package model provide the data structure of all models
package model

import "time"
`

var StructTemplate = `
type {{.Name}} struct {
	ID int `+"`"+`json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id" swaggerignore:"true"`+"`"+` {{range .Columns}}
	{{.Name}} {{.Type}} `+"`"+`json:"{{.LowerCamelCase}},omitempty" form:"{{.LowerCamelCase}}"`+"`"+`{{end}}    
	CreatedAt time.Time `+"`"+`json:"-" swaggerignore:"true"`+"`"+`
	UpdatedAt time.Time `+"`"+`json:"-" swaggerignore:"true"`+"`"+` 
}

type {{.Name}}Param struct {
	*{{.Name}}
	Limit     int  `+"`"+`json:"limit" form:"limit"`+"`"+`
	Offset    int  `+"`"+`json:"offset" form:"offset"`+"`"+`
	NeedCount bool `+"`"+`json:"needCount" form:"needCount"`+"`"+`
}

// TableName will use the name of the table for gorm
func ({{.Name}}) TableName() string {
	return "{{.SnakeCase}}"
}
`
