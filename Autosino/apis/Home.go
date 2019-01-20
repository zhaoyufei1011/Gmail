package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context)  {
	c.String(http.StatusOK,"Home")
}