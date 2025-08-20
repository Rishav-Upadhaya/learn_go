package handlers

import (
	"learn_golang/models"
	"learn_golang/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	myCtx := ctx.Request.Context()
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   err.Error(),
			"message": "Invalid input data",
		})
		return
	}
	createdUser, err := repositories.SaveUser(myCtx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to create user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    createdUser,
	})
}

func HandleGetAllUsers(ctx *gin.Context) {
	myCtx := ctx.Request.Context()
	users, err := repositories.GetAllUsers(myCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to retrieve users",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Users retrieved successfully",
		"users":   users,
	})
}
