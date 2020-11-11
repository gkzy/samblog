package main

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/samblog/conn"
	"github.com/gkzy/samblog/router"
)

func init() {
	gow.InitConfig()
	conn.InitLog()
	conn.InitMySQL()

}

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	r.Static("/static", "static")
	router.PageRouter(r)
	r.Run()
}
