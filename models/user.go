package models

import (
	"github.com/gkzy/gow/lib/mysql"
	"github.com/gkzy/samblog/enum"
)

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

// TableName
func (*User) TableName() string {
	return "tbl_user"
}

// GetUser
func (m *User) GetUser(username string) error {
	db := mysql.GetORM()
	return db.Model(m).Where("username=?", username).First(&m).Error
}

// GetUserList
func (m *User) GetUserList() (data []*Month, err error) {
	db := mysql.GetORM()
	err = db.Model(m).Find(&data).Error
	return
}
