package models

// Category 分类
type Category struct {
	Cid         int64  `gorm:"primary_key;column:cid" json:"cid"`
	Name        string `gorm:"column:name" json:"name"`
	EName       string `gorm:"column:e_name" json:"e_name"`
	Icon        string `gorm:"column:icon" json:"icon"`
	Count       int    `gorm:"column:count" json:"count"`
	Title       string `gorm:"column:title" json:"title"`
	Keyword     string `gorm:"column:keyword" json:"keyword"`
	Description string `gorm:"column:description" json:"description"`
	DateTime
}

func (*Category) TableName() string {
	return "tbl_category"
}
