// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T00:47:23+08:00
// File: model.go

// Package dal represents as Data access layer
package dal

import (
	"context"
	"crud/model"
	"log"
	"time"
)

// GetTeachers will query Teacher by ids, limit and offset
func GetTeachers(ctx context.Context, ids []string, limit int, offset int) ([]model.Teacher, error) {
	var teachers []model.Teacher
	if len(ids) == 0 {
		return nil, nil
	}
	conn := DB.WithContext(ctx)

	if limit <= 0 && limit >= 500 {
		limit = 100
	}
	conn = conn.Limit(limit)
	if offset > 0 {
		conn = conn.Offset(offset)
	}

	err := conn.Find(&teachers, ids).Error
	if err != nil {
		log.Println("func GetTeachers failed: ", err)
		return nil, err
	}
	return teachers, nil
}

// CreateTeacher will create a(n) Teacher by *model.Teacher
func CreateTeacher(ctx context.Context, Teacher *model.Teacher) error {
	conn := DB.WithContext(ctx)
	Teacher.CreateTime = time.Now()
	Teacher.UpdateTime = time.Now()
	err := conn.Create(&Teacher).Error
	if err != nil {
		log.Println("func CreateTeacher failed: ", err)
		return err
	}
	return nil
}

// UpdateTeacher will update a(n) Teacher by *model.Teacher.ID and set the value to *model.Teacher
func UpdateTeacher(ctx context.Context, Teacher *model.Teacher) error {
	conn := DB.WithContext(ctx)
	Teacher.UpdateTime = time.Now()
	err := conn.Where("id = ?", Teacher.ID).Updates(Teacher).Error
	if err != nil {
		log.Println("func UpdateTeacher failed: ", err)
		return err
	}
	return nil
}

// DeleteTeachers will delete all Teacher by ids
func DeleteTeachers(ctx context.Context, ids []string) ([]model.Teacher, error) {
	var teachers []model.Teacher
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&teachers).Error
	if err != nil {
		log.Println("func DeleteTeachers failed: ", err)
		return nil, err
	}
	return teachers, nil
}
