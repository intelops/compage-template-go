package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kube-tarian/compage-template-go/pkg/models"
	"github.com/kube-tarian/compage-template-go/pkg/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateUser(context *gin.Context) {
	// validate input
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := service.UserService{}

	// trigger user creation
	if err := userService.CreateUser(input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func UpdateUser(context *gin.Context) {
	// validate input
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := context.Param("id")

	userService := service.UserService{}

	// trigger user updation
	if err := userService.UpdateUser(id, input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func GetUser(context *gin.Context) {
	id := context.Param("id")

	userService := service.UserService{}

	// trigger user fetching
	user, err := userService.GetUser(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func DeleteUser(context *gin.Context) {
	id := context.Param("id")

	userService := service.UserService{}

	// trigger user deletion
	if err := userService.DeleteUser(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func ListUsers(context *gin.Context) {
	userService := service.UserService{}

	// trigger all users fetching
	users, err := userService.ListUsers()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}
