package apis

import (
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
	name := c.PostForm("account")
	passwd := c.PostForm("password")
	if name == "root" && passwd == "123456" {
		//	c.String(http.StatusOK, "test")
		c.Redirect(302, "/home")
	}
}
