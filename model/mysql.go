package model

import (
	"devops/log"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// MySQL driver.
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataBase struct {
	MyDB *gorm.DB
}

var DB *DataBase

func (db *DataBase) Init() {
	DB = &DataBase{
		MyDB: GetMySqlDB(),
	}
}

func (db *DataBase) Close() {
	DB.MyDB.Close()
}

func InitSelfDB() *gorm.DB {
	db := openDB(viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.dbname"),
		viper.GetUint("database.port"))
	return db
}

func GetMySqlDB() *gorm.DB {
	return InitSelfDB()
}

func openDB(username, password, host, name string, port uint) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		host,
		port,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf("Database connection failed. Database name: %s,Eroor:%s", name, err.Error())
	}

	setupDB(db)
	db.SingularTable(true)
	return db
}

func setupDB(db *gorm.DB) {
	// 用于设置闲置的连接数.
	db.DB().SetMaxIdleConns(5)
}
