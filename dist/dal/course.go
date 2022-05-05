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

// GetCourses will query Course by ids, limit and offset
func GetCourses(ctx context.Context, ids []string, limit int, offset int) ([]model.Course, error) {
	var courses []model.Course
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

	err := conn.Find(&courses, ids).Error
	if err != nil {
		log.Println("func GetCourses failed: ", err)
		return nil, err
	}
	return courses, nil
}

// CreateCourse will create a(n) Course by *model.Course
func CreateCourse(ctx context.Context, course *model.Course) error {
	conn := DB.WithContext(ctx)
	err := conn.Create(&course).Error
	if err != nil {
		log.Println("func CreateCourse failed: ", err)
		return err
	}
	return nil
}

// UpdateCourse will update a(n) Course by *model.Course.ID and set the value to *model.Course
func UpdateCourse(ctx context.Context, course *model.Course) error {
	conn := DB.WithContext(ctx)
	err := conn.Where("id = ?", course.ID).Updates(course).Error
	if err != nil {
		log.Println("func UpdateCourse failed: ", err)
		return err
	}
	return nil
}

// DeleteCourses will delete all Course by ids
func DeleteCourses(ctx context.Context, ids []string) ([]model.Course, error) {
	var courses []model.Course
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&courses).Error
	if err != nil {
		log.Println("func DeleteCourses failed: ", err)
		return nil, err
	}
	return courses, nil
}
