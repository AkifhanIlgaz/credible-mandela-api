package response

import "github.com/gin-gonic/gin"

const (
	StatusSuccess string = "success"
	StatusError   string = "error"
)

func WithSuccess(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, data)
}

func WithError(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status":  StatusError,
		"message": message,
	})
}
