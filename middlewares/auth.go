package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/skorlife/sk-blog/pkg/helper/auth"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, "Invalid Token")
			context.Abort()
			return
		}
		context.Next()
	}
}
