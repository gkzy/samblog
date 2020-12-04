package models

import "github.com/gkzy/gow/lib/mysql"

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

// GetCategoryList
func (m *Category) GetCategoryList() (data []*Category, err error) {
	db := mysql.GetORM()
	err = db.Model(m).Find(&data).Error
	return
}

// Create
func (m *Category) Create() error {
	db := mysql.GetORM()
	return db.Model(m).Create(&m).Error
}

// Delete
func (m *Category) Delete(cid int64) error {
	db := mysql.GetORM()
	return db.Model(m).Where("cid=?", cid).Delete(&m).Error
}

// Get
func (m *Category) Get(cid int64) error {
	db := mysql.GetORM()
	return db.Model(m).Where("cid=?", cid).First(&m).Error
}

// Update
func (m *Category) Update(mp map[string]interface{}) error {
	db := mysql.GetORM()
	return db.Model(m).Updates(mp).Error
}
