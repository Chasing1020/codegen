// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T00:47:23+08:00
// File: response.go

// Package model provide the data structure of all models
package model

// Resp Body of a REST API request
type Resp struct {
	Code    int         `json:"code"`    // Code http.status_code
	Message string      `json:"message"` // Error or success message
	Data    interface{} `json:"data"`    // Response data
}
