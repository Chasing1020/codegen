/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:18
File: user.go
*/

package model

import "time"

type User struct {
	ID         string    `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true" form:"id"`
	Name       string    `json:"name,omitempty" form:"name"`
	Password   string    `json:"password,omitempty" form:"password"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}

func (User) TableName() string {
	return "user"
}
