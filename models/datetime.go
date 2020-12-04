package models

// DateTime datetime struct
type DateTime struct {
	Updated    int    `gorm:"column:updated" json:"-"`
	Created    int    `gorm:"column:created" json:"-"`
	UpdatedStr string `gorm:"-" json:"updated_str"`
	CreatedStr string `gorm:"-" json:"created_str"`
}
