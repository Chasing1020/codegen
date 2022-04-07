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

func GetBooks(ctx context.Context, ids []string, limit int, offset int) ([]model.Book, error) {
	var books []model.Book
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

	err := conn.Find(&books, ids).Error
	if err != nil {
		log.Fatal("func GetBooks failed: ", err)
		return nil, err
	}
	return books, nil
}

func CreateBook(ctx context.Context, Book *model.Book) error {
	conn := DB.WithContext(ctx)
	Book.CreateTime = time.Now()
	Book.UpdateTime = time.Now()
	err := conn.Create(&Book).Error
	if err != nil {
		log.Println("func CreateBooks failed: ", err)
		return err
	}
	return nil
}

func UpdateBook(ctx context.Context, Book *model.Book) error {
	conn := DB.WithContext(ctx)
	Book.UpdateTime = time.Now()
	err := conn.Where("id = ?", Book.ID).Updates(Book).Error
	if err != nil {
		log.Println("func UpdateBook failed: ", err)
		return err
	}
	return nil
}

func DeleteBooks(ctx context.Context, ids []string) ([]model.Book, error) {
	var books []model.Book
	conn := DB.WithContext(ctx)
	err := conn.Where("id IN ?", ids).Delete(&books).Error
	if err != nil {
		log.Println("func DeleteBooks failed: ", err)
		return nil, err
	}
	return books, nil
}