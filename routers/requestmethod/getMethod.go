package requestmethod

import (
	"github.com/gin-gonic/gin"
)

//首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
func getMethod(r *gin.Engine) {
	r.GET("/get/:id", func(c *gin.Context) {
		//将请求中的参数拿出来
		id := c.Param("id")
		user := c.Query("user")
		pwd := c.Query("pwd")


		c.JSON(200, gin.H{
			"id": id,
			"user": user,
			"pwd": pwd,
		})
	})


}
