/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:24
File: response.go
*/

package model

var ResponseTemplate = `/*
Copyright © 2022 {{.Author}} <{{.Email}}>
Time: {{.Time}}
File: response.go
*/

package model

// Resp Body of a REST API request
type Resp struct {
	Code       int         `+"`"+`json:"code"`+"`"+`
	Message    string      `+"`"+`json:"message"`+"`"+`
	Data       interface{} `+"`"+`json:"data"`+"`"+`
}
`
