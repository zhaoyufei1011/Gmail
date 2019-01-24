package apis

import (
	"Autosino/db"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	t, err := template.ParseFiles("view/login.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登陆页"})
}
func Logins(c *gin.Context) {
	var Users db.User
	Users.Name = c.PostForm("account")
	Users.Passwd = c.PostForm("password")
	if Users.Name == "root" && Users.Passwd == "123456" {
		//	c.String(http.StatusOK, "test")
	//	c.String(http.StatusOK, *gin.H{"body": "登陆成功"})
		c.Redirect(302, "/home")
	}
}
