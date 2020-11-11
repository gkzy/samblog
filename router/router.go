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
	r.Any("/user/{month:int}", handler.ArchiveUser)
}
