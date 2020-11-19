/*



/article/{id:int}  # 文章详情

/category  # 所有文章列表
/category/golang  # 分类文章列表

/archive   # 所有月份列表
/archive/202012  # 某个月份的文章列表

/user   # 所有用户列表
/user/1  # 某个用户的文章列表



*/

package router

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/samblog/handler"
	"github.com/gkzy/samblog/middleware"
)

// PageRouter page router
func PageRouter(r *gow.Engine) {

	r.Use(middleware.Version())

	r.Any("/", handler.IndexHandler)
	r.Any("/item/{ename}", handler.ArticleList)
	r.Any("/article/{id:int}", handler.ArticleDetail)
	r.Any("/archive", handler.Archive)
	r.Any("/archive/{month:int}", handler.ArchiveMonth)
	r.Any("/user/{uid:int}", handler.ArchiveUser)
}
