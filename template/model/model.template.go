/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-18:13
File: model.go
*/

package model

var HeadTemplate = `/*
Copyright © 2022 {{.Author}} <{{.Email}}>
Time: {{.Time}}
File: response.go
*/

package model

import "time"
`

var StructTemplate = `
type {{.Name}} struct {
	ID         int `+"`"+`json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id"`+"`"+` {{range .Columns}}
	{{.Name}}  {{.Type}} `+"`"+`json:"{{.Tag}},omitempty" form:"{{.Tag}}"`+"`"+`{{end}}    
	CreateTime time.Time `+"`"+`json:"createTime,omitempty"`+"`"+`
	UpdateTime time.Time `+"`"+`json:"updateTime,omitempty"`+"`"+`
}

func ({{.Name}}) TableName() string {
	return "{{.Tag}}"
}
`
