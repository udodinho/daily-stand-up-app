package response

import "github.com/gin-gonic/gin"

func Failure(code int, c *gin.Context, message string, err bool) {
	c.JSON(code, gin.H{
		"message": message,
		"error":   err,
	})
}
