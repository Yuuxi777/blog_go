package result

import "github.com/gin-gonic/gin"

func Success(msg string, data ...interface{}) gin.H {
	result := gin.H{
		"code":   0,
		"status": 0,
		"msg":    msg,
		"data":   data,
	}
	return result
}

func Fail(code int, msg string) gin.H {
	result := gin.H{
		"code":   code,
		"status": -1,
		"msg":    msg,
	}
	return result
}
