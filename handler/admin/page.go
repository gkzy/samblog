/*
admin page
 */
package admin

import gow "github.com/gkzy/gov"

// Login 登录
func Login(c *gow.Context) {

}

// Logout 退出
func Logout(c *gow.Context) {

}

// Main
func Main(c *gow.Context) {

}

func Article(c *gow.Context) {
	c.HTML("admin/article.html")
}

func ArticleAdd(c *gow.Context) {
	c.HTML("admin/article_add.html")
}
