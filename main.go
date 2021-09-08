package main

import (
	"devops/config"
	"devops/handler"
	"devops/log"
	"devops/middleware"
	"devops/model"
	"devops/router"
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	cfg = flag.String("config", "", "")
)

func init() {
	handler.InitViper()
	handler.InitDB()
	handler.InitHandler()
}

func main() {
	flag.Parse()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	model.DB.Init()
	defer model.DB.Close()
	r := gin.New()
	router.Load(
		r,
		middleware.ProcessTraceID(),
		middleware.Logging(),
	)
	port := viper.GetString("addr")
	log.Log.Info("开始监听http地址", port)
	log.Log.Info(http.ListenAndServe(port, r).Error())
}
