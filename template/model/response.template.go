/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:24
File: response.go
*/

package model

var ResponseTemplate = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{.Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: response.go

// Package model provide the data structure of all models
package model

// Resp Body of a REST API request
type Resp struct {
	Code       int         `+"`"+`json:"code"`+"`"+` // Code http.status_code
	Message    string      `+"`"+`json:"message"`+"`"+` // Error or success message
	Data       interface{} `+"`"+`json:"data"`+"`"+` // Response data
}
`
