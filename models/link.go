package models

import "github.com/gkzy/gow/lib/mysql"

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

// Create
func (m *Link) Create() error {
	db := mysql.GetORM()
	return db.Model(m).Create(&m).Error
}

// Delete
func (m *Link) Delete(id int) error {
	db := mysql.GetORM()
	return db.Model(m).Where("id=?", id).Delete(&m).Error
}

// Get
func (m *Link) Get(id int) error {
	db := mysql.GetORM()
	return db.Model(m).Where("id=?", id).First(&m).Error
}

// Update
func (m *Link) Update(mp map[string]interface{}) error {
	db := mysql.GetORM()
	return db.Model(m).Updates(mp).Error
}

// GetLinkList
func (m *Link) GetLinkList() (data []*Link, err error) {
	db := mysql.GetORM()
	err = db.Model(m).Find(&data).Error
	return
}
