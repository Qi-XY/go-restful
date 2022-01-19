package requestmethod

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {
	getMethod(e)
	postMethod(e)
	putMethod(e)
	deleteMethod(e)
}
