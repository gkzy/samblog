package models

// Blog 信息
type Blog struct {
	Id          int    `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Keyword     string `gorm:"column:keyword" json:"keyword"`
	Description string `gorm:"column:description" json:"description"`
	RecordNum   string `gorm:"column:record_num" json:"record_num"`
	DateTime
}
