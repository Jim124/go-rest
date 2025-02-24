package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPathId(context *gin.Context) {
	id, error := strconv.ParseInt(context.Param("id"), 10, 64)
	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "could not get id from url"})
		return
	}
	context.Set("eventId", id)
	context.Next()
}
