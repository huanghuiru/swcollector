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

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Hostinfo struct {
	ID     int64  `json:"id" `
	IP   string `json:"ip"`
}

func GetIp() (ips []string, err error) {
	portal, err := gorm.Open("mysql", "root:root@tcp(10.112.95.1:3306)/swcollector")
	if err != nil {
		err = fmt.Errorf("connect to swcollector: %s", err.Error())
		portal.Close()
		return
	}
	portal.SingularTable(true)
	portal.AutoMigrate(&Hostinfo{})
	var hostinfo Hostinfo
	dt := portal.First(&hostinfo)
	portal.Close()
	if dt.Error != nil {
		err = dt.Error
		return
	}
	ips = []string{hostinfo.IP}
	return
}

func (this Hostinfo) TableName() string {
	return "hostinfo"
}
