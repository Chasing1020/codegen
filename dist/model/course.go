// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: response.go

// Package model provide the data structure of all models
package model

import "time"

type Course struct {
	ID           int       `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id" swaggerignore:"true"`
	CourseId     string    `json:"course_id,omitempty" form:"course_id"`
	Name         string    `json:"name,omitempty" form:"name"`
	Credit       int       `json:"credit,omitempty" form:"credit"`
	Time         int       `json:"time,omitempty" form:"time"`
	DepartmentId string    `json:"department_id,omitempty" form:"department_id"`
	CreatedAt    time.Time `json:"create_at,omitempty" swaggerignore:"true"`
	UpdatedAt    time.Time `json:"update_at,omitempty" swaggerignore:"true"`
}

// TableName will use the name of the table for gorm
func (Course) TableName() string {
	return "course"
}
