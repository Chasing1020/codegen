/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-04-07 23:11:10.330178 +0800 CST m=+0.001950585
File: model.go
*/

package dal

import (
	"crud/model"
	"context"
	"log"
	"time"
)

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
		log.Fatal("func GetTeachers failed: ", err)
		return nil, err
	}
	return teachers, nil
}

func CreateTeacher(ctx context.Context, Teacher *model.Teacher) error {
	conn := DB.WithContext(ctx)
	Teacher.CreateTime = time.Now()
	Teacher.UpdateTime = time.Now()
	err := conn.Create(&Teacher).Error
	if err != nil {
		log.Println("func CreateTeachers failed: ", err)
		return err
	}
	return nil
}

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