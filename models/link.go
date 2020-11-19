package models

// Link 友情链接
type Link struct {
	Id   int    `gorm:"primary_key;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Url  string `gorm:"column:url" json:"url"`
	DateTime
}

// TableName
func (*Link) TableName() string {
	return "tbl_link"
}
