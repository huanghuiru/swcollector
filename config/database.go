package config

import (
"log"
"fmt"

_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)

type DBPool struct {
	Switchboard    *gorm.DB
}

type Equipment struct {
	ID          int64   `json:"id" `
	Type	    string	`json:"type" `
	Hostname	string	`json:"hostname" `
	Ipaddr		string	`json:"ipaddr" `
	Sn			string	`json:"sn" `
	Os			string	`json:"os" `
	Site		string	`json:"site" `
	Location	string	`json:"location" `
	Model		string	`json:"model" `
	Description	string	`json:"description" `
	Password	string	`json:"password" `
	Nodegroup	string	`json:"nodegroup" `
	Enable		bool	`json:"enable" `

}

var (
	dbp DBPool
	swinfos []Equipment
)

func Con() DBPool {
	return dbp
}

func Info() []Equipment {
	return swinfos[0:3]
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

func GetInfo() (err error) {
	db := Con().Switchboard
	db.AutoMigrate(&Equipment{})
	equipment := []Equipment{}
	dt := db.Find(&equipment)
	if dt.Error != nil {
		err = dt.Error
		log.Println(err)
		return
	}
	swinfos = equipment
	return
}

func GetPassword(swinfos []Equipment,ip string) (pw string,err error) {
	if len(swinfos) > 0 {
		for _, swinfo := range swinfos {
			if swinfo.Ipaddr == ip {
				pw = swinfo.Password
				return
			}
		}
		err = fmt.Errorf("%s don't has snmp password",ip)
		log.Println(err)
		return
	}
	err = fmt.Errorf("there is no switchboard info")
	log.Println(err)
	return
}

func (this Equipment) TableName() string {
	return "equipment"
}

func CloseDB() (err error) {
	err = dbp.Switchboard.Close()
	if err != nil {
		return
	}
	return
}
