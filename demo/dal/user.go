/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:45
File: user.go
*/

package dal

import (
	"codegen/model"
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
		log.Println("func `GetUsers` failed: ", err)
		return nil, err
	}
	return users, nil
}

func CreateUser(ctx context.Context, user *model.User) error {
	conn := DB.WithContext(ctx)
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	err := conn.Create(&user).Error
	if err != nil {
		log.Println("func `CreateUser` failed: ", err)
		return err
	}
	return nil
}

func UpdateUser(ctx context.Context, user *model.User) error {
	conn := DB.WithContext(ctx)
	user.UpdateTime = time.Now()
	err := conn.Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		log.Println("func `UpdateUser` failed: ", err)
		return err
	}
	return nil
}

func DeleteUsers(ctx context.Context, ids []string) ([]model.User, error) {
	var users []model.User
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&users).Error
	if err != nil {
		log.Println("func `DeleteUsers` failed: ", err)
		return nil, err
	}
	return users, nil
}
