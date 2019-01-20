package router

import(
	"github.com/gin-gonic/gin"
	"Autosino/apis"
)
func InitRouter()*gin.Engine  {
	router:=gin.Default()
	router.GET("/",apis.Login)
	return router
}