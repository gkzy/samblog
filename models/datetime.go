package models

// DateTime datetime struct
type DateTime struct {
	Deleted    int    `gorm:"column:deleted" json:"-"`
	Updated    int    `gorm:"column:updated" json:"-"`
	Created    int    `gorm:"column:created" json:"-"`
	DeletedStr string `gorm:"-" json:"deleted_str"`
	UpdatedStr string `gorm:"-" json:"updated_str"`
	CreatedStr string `gorm:"-" json:"created_str"`
}
