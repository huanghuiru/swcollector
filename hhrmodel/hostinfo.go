// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hhrmodel

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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

func GetInfo() (swinfos []Equipment, err error) {
	portal, err := gorm.Open("mysql", "root:root@tcp(10.112.95.1:3306)/gnet")
	defer portal.Close()
	if err != nil {
		err = fmt.Errorf("connect to swcollector: %s", err.Error())
		log.Println(err)
		return
	}
	portal.SingularTable(true)
	portal.AutoMigrate(&Equipment{})
	var equipment Equipment
	dt := portal.First(&equipment)
	if dt.Error != nil {
		err = dt.Error
		log.Println(err)
		return
	}
	swinfos = []Equipment{equipment}
	return
}

func GetPassword(swinfos []Equipment,ip string) (pw string,err error) {
	if len(swinfos) > 0 {
		for _, swinfo := range swinfos {
			if swinfo.Ipaddr == ip {
				pw = swinfo.Password
				return
			}
			err = fmt.Errorf("%s don't has snmp password",ip)
			log.Println(err)
			return
		}
	}
}

func (this Equipment) TableName() string {
	return "equipment"
}
