// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: response.go

// Package model provide the data structure of all models
package model

import "time"

type Elective struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id" swaggerignore:"true"`
	StudentId string    `json:"student_id,omitempty" form:"student_id"`
	Term      string    `json:"term,omitempty" form:"term"`
	CourseId  string    `json:"course_id,omitempty" form:"course_id"`
	TeacherId string    `json:"teacher_id,omitempty" form:"teacher_id"`
	Score     string    `json:"score,omitempty" form:"score"`
	CreatedAt time.Time `json:"create_at,omitempty" swaggerignore:"true"`
	UpdatedAt time.Time `json:"update_at,omitempty" swaggerignore:"true"`
}

// TableName will use the name of the table for gorm
func (Elective) TableName() string {
	return "elective"
}
