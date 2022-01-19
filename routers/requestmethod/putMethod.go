package requestmethod

import (
	"github.com/gin-gonic/gin"
)

func putMethod(r *gin.Engine) {

	r.PUT("/put/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.PostForm("user")
		pwd := c.PostForm("pwd")

		c.JSON(200, gin.H{
			"id": id,
			"user": user,
			"pwd": pwd,
		})
	})

}
