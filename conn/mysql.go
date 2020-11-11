package conn

import (
	"fmt"
	"github.com/gkzy/gow/lib/config"
	"github.com/gkzy/gow/lib/logy"
	"github.com/gkzy/gow/lib/mysql"
)

func InitMySQL() {
	name := config.GetString("defaultDB")
	conf := &mysql.DBConfig{
		Name:     name,
		User:     config.GetString(name + "::user"),
		Password: config.GetString(name + "::password"),
		Host:     config.GetString(name + "::host"),
		Port:     config.DefaultInt(name+"::port", 3306),
	}

	runMode := config.DefaultString("run_mode", "dev")
	if runMode == "dev" {
		conf.Debug = true
	}

	err := mysql.InitDefaultDB(conf)
	if err != nil {
		logy.Panic(err)
	}
	logy.Info(fmt.Sprintf("[DB]-[%v] initialized successfully => %v", conf.Name, conf))
}
