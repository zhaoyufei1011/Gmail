package db

import (
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	yaml "gopkg.in/yaml.v2"
)

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var DB *gorm.DB

func init() {
	var err error
	var conf Database
	confg := conf.Getconf()
	sdn := confg.Username + ":" + confg.Password + "@tcp(" + confg.Host + ":" + confg.Port + ")/" + confg.Database + "?charset=utf8&&parseTime=true"
	DB, err = gorm.Open("mysql", sdn)
	if err != nil {
		fmt.Println(err)
	}
	DB.DB().SetMaxIdleConns(100)
	DB.DB().SetMaxOpenConns(10)
	DB.SingularTable(true)
}

func (conf *Database) Getconf() *Database {
	yamlFile, err := ioutil.ReadFile("./conf/db.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
