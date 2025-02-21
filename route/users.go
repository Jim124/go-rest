package route

import (
	"fmt"
	"go-rest/models"
	"go-rest/utils"
	"log"
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
	fmt.Println(user.ID)
	if error != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}
	token, error := utils.GenerateToken(user.ID, user.Email)
	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		log.Fatal(error)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successfully", "token": token})
}
