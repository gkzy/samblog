package models

// Month 分份
type Month struct {
	Id    int64 `gorm:"primary_key;column:id" json:"id"`
	Month int   `gorm:"column:month" json:"month"`
	Count int   `gorm:"column:count" json:"count"`
	DateTime
}

func (*Month) TableName() string {
	return "tbl_month"
}
