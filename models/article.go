package models

import (
	"github.com/gkzy/gow/lib/mysql"
	"github.com/gkzy/samblog/enum"
)

// Article 文章
type Article struct {
	Aid          int64            `json:"aid" gorm:"primary_key;column:aid"`         //id
	Cid          int64            `json:"cid" gorm:"column:cid"`                     //分类id
	Title        string           `json:"title" gorm:"column:title"`                 //标题
	Photo        string           `json:"photo" gorm:"column:photo"`                 //图片
	STitle       string           `json:"s_title" gorm:"column:s_title"`             //短标题
	SContent     string           `json:"s_content" gorm:"column:s_content"`         //内容简介 200字内
	Author       string           `json:"author" gorm:"column:author"`               //作者
	Uid          int64            `json:"uid" gorm:"column:uid"`                     //用户uid
	Username     string           `json:"username" gorm:"column:username"`           //用户名
	Views        int              `json:"views" gorm:"column:views"`                 //点击量
	IsTop        int              `json:"is_top" gorm:"column:is_top"`               //是否置顶
	IsHot        int              `json:"is_hot" gorm:"column:is_hot"`               //是否热门
	IsCopyright  int              `json:"is_copyright" gorm:"column:is_copyright"`   //是否有版权信息
	Password     string           `json:"password" gorm:"column:password"`           //密码
	IP           string           `json:"ip" gorm:"column:ip"`                       // IP
	Month        int              `json:"month" gorm:"column:month"`                 //202011
	ArticleState enum.CommonState `json:"article_state" gorm:"column:article_state"` //文章状态
	ArticleType  enum.ArticleType `json:"article_type" gorm:"column:article_type"`   //状态类型
	DateTime
}

func (*Article) TableName() string {
	return "tbl_article"
}

// Get
func (m *Article) Get(aid int64) error {
	db := mysql.GetORM()
	return db.Model(m).Where("aid=?", aid).First(&m).Error
}

// Page 翻页列表
func (m *Article) Page(cid, uid, month int64, articleState, articleType int, offset, limit int64) (data []*Article, count int64, err error) {
	db := mysql.GetORM()
	if cid > 0 {
		db = db.Where("cid=?", cid)
	}
	if uid > 0 {
		db = db.Where("uid=?", uid)
	}
	if month > 0 {
		db = db.Where("month=?", month)
	}
	if articleState != -100 {
		db = db.Where("article_state=?", articleState)
	}
	if articleType != -100 {
		db = db.Where("article_type=?", articleType)
	}

	err = db.Order("aid desc").Offset(offset).Limit(limit).Find(&data).Error
	err = db.Count(&count).Error
	return
}

// Create create article
//	insert article and article_detail
func (m *Article) Create(content, mdContent string) error {
	g := mysql.GetORM().Begin()

	err := g.Create(&m).Error
	if err != nil {
		g.Rollback()
		return err
	}

	detail := &ArticleDetail{
		Aid:       m.Aid,
		Content:   content,
		MDContent: mdContent,
	}

	err = g.Create(&detail).Error
	if err != nil {
		g.Rollback()
		return err
	}
	return g.Commit().Error

}

// Delete
func (m *Article) Delete(aid int64) error {
	db := mysql.GetORM()
	return db.Model(m).Where("aid=?", aid).Delete(&m).Error
}

// Update
func (m *Article) Update(mp map[string]interface{}) error {
	db := mysql.GetORM()
	return db.Model(m).Updates(mp).Error
}

// ArticleDetail 文章详情
type ArticleDetail struct {
	Aid       int64  `json:"aid" gorm:"column:aid"`               //id
	Content   string `json:"content" gorm:"column:content"`       //内容
	MDContent string `json:"md_content" gorm:"column:md_content"` //markdown 内容
}

func (*ArticleDetail) TableName() string {
	return "tbl_article_detail"
}

// Get
func (m *ArticleDetail) Get(aid int64) error {
	db := mysql.GetORM()
	return db.Model(m).Where("aid=?", aid).First(&m).Error
}
