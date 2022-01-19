package requestmethod

import (
	"github.com/gin-gonic/gin"
)

func postMethod(r *gin.Engine) {
	r.POST("/post", func(c *gin.Context) {
		user := c.PostForm("user")
		pwd := c.PostForm("pwd")

		c.JSON(200, gin.H{
			"user": user,
			"pwd": pwd,
		})
	})







}
