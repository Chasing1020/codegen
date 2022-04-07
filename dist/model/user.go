/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-04-07 23:11:10.330178 +0800 CST m=+0.001950585
File: response.go
*/

package model

import "time"

type User struct {
	ID         int `json:"id,omitempty" gorm:"primaryKey;autoIncrement" form:"id"` 
	Name  string `json:"name,omitempty" form:"name"`
	Password  string `json:"password,omitempty" form:"password"`
	Grade  string `json:"grade,omitempty" form:"grade"`    
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}

func (User) TableName() string {
	return "user"
}
