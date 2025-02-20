package route

import (
	"go-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	error := context.ShouldBindBodyWithJSON(&user)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	error = user.Save()
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "created user successfully"})
}

func login(context *gin.Context) {
	var user models.User
	error := context.ShouldBindBodyWithJSON(&user)
	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	error = user.GetUserByEmail()
	if error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successfully"})
}
