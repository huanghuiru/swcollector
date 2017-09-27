package config

import (
"log
"fmt"

_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)

type DBPool struct {
	Switchboard    *gorm.DB
}

var (
	dbp DBPool
)

func Con() DBPool {
	return dbp
}

func InitDB() (err error) {
	portal, err := gorm.Open("mysql", "root:root@tcp(10.112.95.1:3306)/gnet")
	if err != nil {
		err = fmt.Errorf("connect to swcollector: %s", err.Error())
		log.Println(err)
		return
	}
	portal.SingularTable(true)

	dbp.Switchboard = portal

	return
}

func CloseDB() (err error) {
	err = dbp.Switchboard.Close()
	if err != nil {
		return
	}
	return
}
