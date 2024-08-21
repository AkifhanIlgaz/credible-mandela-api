package response

import "github.com/gin-gonic/gin"

const (
	statusSuccess string = "success"
	statusError   string = "error"
)

func WithSuccess(ctx *gin.Context, statusCode int, message string, data any) {
	ctx.JSON(statusCode, gin.H{
		"status":  statusSuccess,
		"message": message,
		"result":  data,
	})
}

func WithError(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status":  statusError,
		"message": message,
	})
}
