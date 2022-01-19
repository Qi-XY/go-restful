package requestmethod

import "github.com/gin-gonic/gin"

func deleteMethod(r *gin.Engine) {

	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})


}