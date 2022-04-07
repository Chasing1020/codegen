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
		log.Fatal("func GetUsers failed: ", err)
		return nil, err
	}
	return users, nil
}

func CreateUser(ctx context.Context, User *model.User) error {
	conn := DB.WithContext(ctx)
	User.CreateTime = time.Now()
	User.UpdateTime = time.Now()
	err := conn.Create(&User).Error
	if err != nil {
		log.Println("func CreateUsers failed: ", err)
		return err
	}
	return nil
}

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