package models

import "github.com/gkzy/samblog/enum"

// User 用户
type User struct {
	Uid       int64            `json:"uid" gorm:"primary_key;column:uid"`
	Username  string           `json:"username" gorm:"column:username"`
	Avatar    string           `json:"avatar" gorm:"column:avatar"`
	Password  string           `json:"password" gorm:"column:password"`
	Sex       enum.SexType     `json:"sex" gorm:"column:sex"`
	Count     int              `json:"count" gorm:"column:count"`
	UserState enum.CommonState `json:"user_state" gorm:"column:user_state"`
	DateTime
}
