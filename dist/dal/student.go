// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T18:13:46+08:00
// File: model.go

// Package dal represents as Data access layer
package dal

import (
	"context"
	"crud/model"
	"log"
)

// GetStudents will query Student by ids, limit and offset
func GetStudents(ctx context.Context, ids []string, limit int, offset int) ([]model.Student, error) {
	var students []model.Student
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

	err := conn.Find(&students, ids).Error
	if err != nil {
		log.Println("func GetStudents failed: ", err)
		return nil, err
	}
	return students, nil
}

// CreateStudent will create a(n) Student by *model.Student
func CreateStudent(ctx context.Context, student *model.Student) error {
	conn := DB.WithContext(ctx)
	err := conn.Create(&student).Error
	if err != nil {
		log.Println("func CreateStudent failed: ", err)
		return err
	}
	return nil
}

// UpdateStudent will update a(n) Student by *model.Student.ID and set the value to *model.Student
func UpdateStudent(ctx context.Context, student *model.Student) error {
	conn := DB.WithContext(ctx)
	err := conn.Where("id = ?", student.ID).Updates(student).Error
	if err != nil {
		log.Println("func UpdateStudent failed: ", err)
		return err
	}
	return nil
}

// DeleteStudents will delete all Student by ids
func DeleteStudents(ctx context.Context, ids []string) ([]model.Student, error) {
	var students []model.Student
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&students).Error
	if err != nil {
		log.Println("func DeleteStudents failed: ", err)
		return nil, err
	}
	return students, nil
}
