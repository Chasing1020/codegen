/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-04-07 23:11:10.330178 +0800 CST m=+0.001950585
File: response.go
*/

package model

import "time"

type Book struct {
	ID         int `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id"` 
	Name  string `json:"name,omitempty" form:"name"`
	Author  string `json:"author,omitempty" form:"author"`
	Price  int `json:"price,omitempty" form:"price"`
	Isbn  string `json:"isbn,omitempty" form:"isbn"`    
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}

func (Book) TableName() string {
	return "book"
}
