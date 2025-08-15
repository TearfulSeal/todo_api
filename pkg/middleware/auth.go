package middleware

import (
	"strings"
	"todo_api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(401, gin.H{"error":"missing token"})
			c.Abort()
			return 
		}
		tokenStr := strings.TrimPrefix(authHeader,"Bearer ")
		userID, err := jwt.ParseToken(tokenStr)
		if err != nil{
			c.JSON(401, gin.H{"error":"invalid token"})
			c.Abort()
			return 
		}
		c.Set("userID",userID)
		c.Next()

	}
}
