package router

import (
	"Autosino/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.Static("/static", "./static")
	//router.StaticFS("/view", http.Dir("view"))
	//	router.Static("/view", root)
	router.GET("/login", apis.Login)
	router.POST("/login", apis.Logins)
	router.GET("/home", apis.Home)
	return router
}
