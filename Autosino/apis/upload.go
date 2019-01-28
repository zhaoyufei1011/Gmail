package apis

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Upfilehtml(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{"title": "文件上传"})
}
func Uploadfile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		panic(err)
	}
	filename := header.Filename
	out, err := os.Create("static/upload/" + filename)
	if err != nil {
		panic(err)
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "", "data": ""})
	/*file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	c.SaveUploadedFile(file, "./static/upload")*/
}
