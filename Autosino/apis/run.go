package apis

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	c.HTML(http.StatusOK, "run.html", gin.H{"title": "运行命令"})
}
func RunCommond(c *gin.Context) {
	run := c.Request.FormValue("run")
	if run {
		command := "php " + c.Request.FormValue("file")
		cmd := exec.Command("/bin/bash", command)
	}

}
