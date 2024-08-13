package response

import "github.com/gin-gonic/gin"

const (
	StatusSuccess string = "success"
	StatusError   string = "error"
)

var (
	successResponseBase = gin.H{
		"status": StatusSuccess,
	}
)

func Success(ctx *gin.Context, statusCode int, data gin.H) {
	ctx.JSON(statusCode, mergeResponseData(successResponseBase, data))
}

func Error(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status":  StatusError,
		"message": message,
	})
}

func mergeResponseData(base gin.H, data gin.H) gin.H {
	for key, value := range data {
		base[key] = value
	}
	return base
}
