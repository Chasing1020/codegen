// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: response.go

// Package model provide the data structure of all models
package model

import "time"

type Department struct {
	ID           int       `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id" swaggerignore:"true"`
	DepartmentId string    `json:"department_id,omitempty" form:"department_id"`
	Name         string    `json:"name,omitempty" form:"name"`
	Address      string    `json:"address,omitempty" form:"address"`
	Phone        string    `json:"phone,omitempty" form:"phone"`
	CreatedAt    time.Time `json:"create_at,omitempty" swaggerignore:"true"`
	UpdatedAt    time.Time `json:"update_at,omitempty" swaggerignore:"true"`
}

// TableName will use the name of the table for gorm
func (Department) TableName() string {
	return "department"
}
