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

// GetElectives will query Elective by ids, limit and offset
func GetElectives(ctx context.Context, ids []string, limit int, offset int) ([]model.Elective, error) {
	var electives []model.Elective
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

	err := conn.Find(&electives, ids).Error
	if err != nil {
		log.Println("func GetElectives failed: ", err)
		return nil, err
	}
	return electives, nil
}

// CreateElective will create a(n) Elective by *model.Elective
func CreateElective(ctx context.Context, elective *model.Elective) error {
	conn := DB.WithContext(ctx)
	err := conn.Create(&elective).Error
	if err != nil {
		log.Println("func CreateElective failed: ", err)
		return err
	}
	return nil
}

// UpdateElective will update a(n) Elective by *model.Elective.ID and set the value to *model.Elective
func UpdateElective(ctx context.Context, elective *model.Elective) error {
	conn := DB.WithContext(ctx)
	err := conn.Where("id = ?", elective.ID).Updates(elective).Error
	if err != nil {
		log.Println("func UpdateElective failed: ", err)
		return err
	}
	return nil
}

// DeleteElectives will delete all Elective by ids
func DeleteElectives(ctx context.Context, ids []string) ([]model.Elective, error) {
	var electives []model.Elective
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&electives).Error
	if err != nil {
		log.Println("func DeleteElectives failed: ", err)
		return nil, err
	}
	return electives, nil
}
