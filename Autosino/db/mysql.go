package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var sdn = "ss数据库连接串"

func init() {
	var err error
	DB, err := gorm.Open("mysql", sdn)
	if err != nil {
		fmt.Println(err)
	}
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetMaxOpenConns(10)
	DB.SingularTable(true)
}
