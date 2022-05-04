// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T00:47:23+08:00
// File: response.go

// Package model provide the data structure of all models
package model

import "time"

type Teacher struct {
	ID         int       `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id" swaggerignore:"true"`
	Name       string    `json:"name,omitempty" form:"name"`
	Department string    `json:"department,omitempty" form:"department"`
	CreateTime time.Time `json:"createTime,omitempty" swaggerignore:"true"`
	UpdateTime time.Time `json:"updateTime,omitempty" swaggerignore:"true"`
}

// TableName will use the name of the table for gorm
func (Teacher) TableName() string {
	return "teacher"
}
