package models

// Tag
type Tag struct {
	Id    int64  `gorm:"primary_key;column:id" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Count int    `gorm:"column:count" json:"count"`
	DateTime
}
