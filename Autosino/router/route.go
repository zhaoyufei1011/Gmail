package router

import (
	"Autosino/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("view/*")

	router.GET("/", apis.Upfilehtml)
	return router
}
