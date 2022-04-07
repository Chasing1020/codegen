/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-17:23
File: response.go
*/

package model

// Resp Body of a REST API request
type Resp struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
	Data       interface{} `json:"data"`
}

