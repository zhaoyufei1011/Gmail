package apis

import (
	"Autosino/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var User db.User //声明数据结构，将web内容存储至结构
	User.Name = c.GetString("name")
	User.Passwd = c.GetString("password")
	db.DB.Save(&User)
}
func Result(c *gin.Context) {
	//	c.JSON(http.StatusOK, Result)
	c.String(http.StatusOK, "jieguo ")
}
