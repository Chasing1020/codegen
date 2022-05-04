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

// GetBooks will query Book by ids, limit and offset
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
		log.Println("func GetBooks failed: ", err)
		return nil, err
	}
	return books, nil
}

// CreateBook will create a(n) Book by *model.Book
func CreateBook(ctx context.Context, Book *model.Book) error {
	conn := DB.WithContext(ctx)
	Book.CreateTime = time.Now()
	Book.UpdateTime = time.Now()
	err := conn.Create(&Book).Error
	if err != nil {
		log.Println("func CreateBook failed: ", err)
		return err
	}
	return nil
}

// UpdateBook will update a(n) Book by *model.Book.ID and set the value to *model.Book
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

// DeleteBooks will delete all Book by ids
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
