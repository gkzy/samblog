package handler

import (
	"github.com/gkzy/gow"
)


// IndexHandler
func IndexHandler(c *gow.Context) {
	c.HTML("index.html")
}

func ArticleList(c *gow.Context) {
	c.HTML("article_list.html")
}

// ArticleDetail
func ArticleDetail(c *gow.Context) {
	c.HTML("article_detail.html")
}


// Archive
func Archive(c *gow.Context){
	c.HTML("archive.html")
}

// ArchiveMonth
func ArchiveMonth(c *gow.Context){
	c.HTML("archive_month.html")
}

// ArchiveUser
func ArchiveUser(c *gow.Context){
	c.HTML("archive_user.html")
}