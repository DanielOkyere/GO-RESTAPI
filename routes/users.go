package routes

import (
	"danielokyere/RESTCRUD/models"
	"danielokyere/RESTCRUD/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	 err = user.Save()
	 if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse request data"})
        return
     }
	 ctx.JSON(http.StatusCreated, gin.H{"message": "User saved successfully", "data": user}) 
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Name, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})

}