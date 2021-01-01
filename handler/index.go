/*
前端handler
*/

package handler

import (
	"github.com/gkzy/gow"
)

// IndexHandler
func Index(c *gow.Context) {
	c.HTML("index.html")
}

// Category 分类目录
//	router /category
func Category(c *gow.Context) {

}

// ArticleList 分类文章列表
//	router /category/golang
func ArticleList(c *gow.Context) {
	c.HTML("article_list.html")
}

// ArticleDetail  文章详情
//	router /article-{id:int}.html
func ArticleDetail(c *gow.Context) {
	c.HTML("article_detail.html")
}

// Archive
//	router /archive
func Archive(c *gow.Context) {
	c.HTML("archive.html")
}

// ArchiveMonth
//	router /archive/202011
func ArchiveMonth(c *gow.Context) {
	c.HTML("archive_month.html")
}

// User
//	router /user
func User(c *gow.Context) {
	c.HTML("user.html")
}

// ArchiveUser
//	router /user/1
func ArchiveUser(c *gow.Context) {
	c.HTML("archive_user.html")
}
