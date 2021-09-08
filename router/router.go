package router

import (
	"devops/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine, middlewares ...gin.HandlerFunc) *gin.Engine {
	engine.Use(gin.Recovery())
	engine.Use(middlewares...)
	engine.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "API路由不正确.")
	})

	account := engine.Group("/v1/account")
	{

		account.GET("", handler.AccountHandle.ListAccount)              // 获取用户列表
		account.GET("/:account_name", handler.AccountHandle.GetAccount) // 获取指定用户的详细信息
		account.POST("", handler.AccountHandle.AccountCreate)           //新增用户
		account.DELETE("/:id", handler.AccountHandle.Delete)            // 删除用户
		account.PUT("/", handler.AccountHandle.Update)                  // 更新用户
		account.POST("/login", handler.AccountHandle.Login)             //登录
	}
	return engine
}
