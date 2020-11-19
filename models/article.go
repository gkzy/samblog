package models

import "github.com/gkzy/samblog/enum"

// Article 文章
type Article struct {
	Aid          int64            `json:"aid" gorm:"primary_key;column:aid"`         //id
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

// ArticleDetail 文章详情
type ArticleDetail struct {
	Aid       int64  `json:"aid" gorm:"column:aid"`               //id
	Content   string `json:"content" gorm:"column:content"`       //内容
	MDContent string `json:"md_content" gorm:"column:md_content"` //markdown 内容
}
