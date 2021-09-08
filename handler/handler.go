package handler

import (
	"devops/conf"
	"devops/config"
	"devops/model"
	"devops/repository"
	"devops/service"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB            *gorm.DB
	AccountHandle AccountHandler
)

func InitViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func InitDB() {
	var err error
	conf := &conf.DBConf{
		Driver:   "",
		Host:     viper.GetString("database.host"),
		Port:     viper.GetUint("database.port"),
		UserName: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.dbname"),
		Charset:  "",
	}

	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.UserName,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DbName,
		true,
		"Local")

	fmt.Println(config)

	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	DB.SingularTable(true)
}

func InitHandler() {
	AccountHandle = AccountHandler{
		Srv: &service.AccountService{
			Repo: &repository.AccountModelRepo{
				DB: model.DataBase{MyDB: DB},
			},
		}}
}
