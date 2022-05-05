// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: response.go

// Package model provide the data structure of all models
package model

import "time"

type Student struct {
	ID           int        `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id" swaggerignore:"true"`
	StudentId    string     `json:"student_id,omitempty" form:"student_id"`
	Password     string     `json:"password,omitempty" form:"password"`
	Name         string     `json:"name,omitempty" form:"name"`
	Sex          string     `json:"sex,omitempty" form:"sex"`
	Birthday     *time.Time `json:"birthday,omitempty" form:"birthday"`
	Hometown     string     `json:"hometown,omitempty" form:"hometown"`
	Phone        string     `json:"phone,omitempty" form:"phone"`
	DepartmentId string     `json:"department_id,omitempty" form:"department_id"`
	CreatedAt    time.Time  `json:"create_at,omitempty" swaggerignore:"true"`
	UpdatedAt    time.Time  `json:"update_at,omitempty" swaggerignore:"true"`
}

// TableName will use the name of the table for gorm
func (Student) TableName() string {
	return "student"
}
