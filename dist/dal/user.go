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

// GetUsers will query User by ids, limit and offset
func GetUsers(ctx context.Context, ids []string, limit int, offset int) ([]model.User, error) {
	var users []model.User
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

	err := conn.Find(&users, ids).Error
	if err != nil {
		log.Println("func GetUsers failed: ", err)
		return nil, err
	}
	return users, nil
}

// CreateUser will create a(n) User by *model.User
func CreateUser(ctx context.Context, User *model.User) error {
	conn := DB.WithContext(ctx)
	User.CreateTime = time.Now()
	User.UpdateTime = time.Now()
	err := conn.Create(&User).Error
	if err != nil {
		log.Println("func CreateUser failed: ", err)
		return err
	}
	return nil
}

// UpdateUser will update a(n) User by *model.User.ID and set the value to *model.User
func UpdateUser(ctx context.Context, User *model.User) error {
	conn := DB.WithContext(ctx)
	User.UpdateTime = time.Now()
	err := conn.Where("id = ?", User.ID).Updates(User).Error
	if err != nil {
		log.Println("func UpdateUser failed: ", err)
		return err
	}
	return nil
}

// DeleteUsers will delete all User by ids
func DeleteUsers(ctx context.Context, ids []string) ([]model.User, error) {
	var users []model.User
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&users).Error
	if err != nil {
		log.Println("func DeleteUsers failed: ", err)
		return nil, err
	}
	return users, nil
}
