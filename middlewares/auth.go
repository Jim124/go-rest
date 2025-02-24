package middlewares

import (
	"go-rest/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	bearerToken := context.Request.Header.Get("Authorization")
	if bearerToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	token := strings.Split(bearerToken, " ")[1]
	userId, error := utils.ValidToken(token)
	if error != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
