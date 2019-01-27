package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upfilehtml(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{"title": "文件上传"})
}
