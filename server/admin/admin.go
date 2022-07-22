package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello administrator")
	re, _ := c.GetPostForm("id")
	fmt.Println(re)
}
func Admin(e *gin.Engine) {
	e.POST("/admin", helloworld)
}
