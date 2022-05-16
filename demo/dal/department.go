// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022-05-05T19:02:15+08:00
// File: model.go

// Package dal represents as Data access layer
package dal

import (
	"context"
	"crud/model"
	"log"
)

// GetDepartments will query Department by ids, limit and offset
func GetDepartments(ctx context.Context, ids []string, limit int, offset int) ([]model.Department, error) {
	var departments []model.Department
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

	err := conn.Find(&departments, ids).Error
	if err != nil {
		log.Println("func GetDepartments failed: ", err)
		return nil, err
	}
	return departments, nil
}

// CreateDepartment will create a(n) Department by *model.Department
func CreateDepartment(ctx context.Context, department *model.Department) error {
	conn := DB.WithContext(ctx)
	err := conn.Create(&department).Error
	if err != nil {
		log.Println("func CreateDepartment failed: ", err)
		return err
	}
	return nil
}

// UpdateDepartment will update a(n) Department by *model.Department.ID and set the value to *model.Department
func UpdateDepartment(ctx context.Context, department *model.Department) error {
	conn := DB.WithContext(ctx)
	err := conn.Where("id = ?", department.ID).Updates(department).Error
	if err != nil {
		log.Println("func UpdateDepartment failed: ", err)
		return err
	}
	return nil
}

// DeleteDepartments will delete all Department by ids
func DeleteDepartments(ctx context.Context, ids []string) ([]model.Department, error) {
	var departments []model.Department
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&departments).Error
	if err != nil {
		log.Println("func DeleteDepartments failed: ", err)
		return nil, err
	}
	return departments, nil
}
