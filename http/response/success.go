package response

import "github.com/gin-gonic/gin"

func Success(code int, data interface{}, c *gin.Context, message string, err bool) {
	c.JSON(code, gin.H{
		"message": message,
		"data":    data,
		"error":   err,
	})
}
