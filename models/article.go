package models

import "github.com/gkzy/samblog/enum"

// Article 文章
type Article struct {
	Aid          int64            `json:"aid" gorm:"primary_key;column:aid"`
	Title        string           `json:"title" gorm:"column:title"`
	Photo        string           `json:"photo" gorm:"column:photo"`
	STitle       string           `json:"s_title" gorm:"column:s_title"`
	SContent     string           `json:"s_content" gorm:"column:s_content"`
	Author       string           `json:"author" gorm:"column:author"`
	Uid          int64            `json:"uid" gorm:"column:uid"`
	Username     string           `json:"username" gorm:"column:username"`
	Views        int              `json:"views" gorm:"column:views"`
	ArticleState enum.CommonState `json:"article_state" gorm:"column:article_state"`
	ArticleType  enum.ArticleType `json:"article_type" gorm:"column:article_type"`
	IsTop        int              `json:"is_top" gorm:"column:is_top"`
	IsHot        int              `json:"is_hot" gorm:"column:is_hot"`
	IsCopyright  int              `json:"is_copyright" gorm:"column:is_copyright"`  //是否有版权信息
	Password     string           `json:"password" gorm:"column:password"`
	IP           string           `json:"ip" gorm:"column:ip"`
	Month        int              `json:"month" gorm:"column:month"`
	DateTime
}

// ArticleDetail 文章详情
type ArticleDetail struct {
	Aid       int64
	Content   string
	MDContent string `json:"md_content" gorm:"column:md_content"`
}
