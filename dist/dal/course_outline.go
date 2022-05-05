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

// GetCourseOutlines will query CourseOutline by ids, limit and offset
func GetCourseOutlines(ctx context.Context, ids []string, limit int, offset int) ([]model.CourseOutline, error) {
	var courseOutlines []model.CourseOutline
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

	err := conn.Find(&courseOutlines, ids).Error
	if err != nil {
		log.Println("func GetCourseOutlines failed: ", err)
		return nil, err
	}
	return courseOutlines, nil
}

// CreateCourseOutline will create a(n) CourseOutline by *model.CourseOutline
func CreateCourseOutline(ctx context.Context, courseOutline *model.CourseOutline) error {
	conn := DB.WithContext(ctx)
	err := conn.Create(&courseOutline).Error
	if err != nil {
		log.Println("func CreateCourseOutline failed: ", err)
		return err
	}
	return nil
}

// UpdateCourseOutline will update a(n) CourseOutline by *model.CourseOutline.ID and set the value to *model.CourseOutline
func UpdateCourseOutline(ctx context.Context, courseOutline *model.CourseOutline) error {
	conn := DB.WithContext(ctx)
	err := conn.Where("id = ?", courseOutline.ID).Updates(courseOutline).Error
	if err != nil {
		log.Println("func UpdateCourseOutline failed: ", err)
		return err
	}
	return nil
}

// DeleteCourseOutlines will delete all CourseOutline by ids
func DeleteCourseOutlines(ctx context.Context, ids []string) ([]model.CourseOutline, error) {
	var courseOutlines []model.CourseOutline
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&courseOutlines).Error
	if err != nil {
		log.Println("func DeleteCourseOutlines failed: ", err)
		return nil, err
	}
	return courseOutlines, nil
}
