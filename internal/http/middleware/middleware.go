package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const requestIDKey string = "requestID"

func SetRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := uuid.New().String()
		ctx.Set(requestIDKey, requestID)
		ctx.Next()
	}
}

func GetRequestID(ctx *gin.Context) string {
	requestID, exists := ctx.Get(requestIDKey)
	if !exists {
		return ""
	}
	return requestID.(string)
}
