package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yatharthx/events-rest-api/models"
	"github.com/yatharthx/events-rest-api/utils"
) 

func createUser(ctx *gin.Context) {
  var user models.User
  err := ctx.ShouldBindJSON(&user)

  if err!=nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message": "Could not parse request data.",
    })
    return
  }

  err = user.Save()

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "message": "Could not save user",
    })
    return
  }

  ctx.JSON(http.StatusCreated, gin.H{
    "message": "User created successfully",
    "user": user,
  })
}

func loginUser(ctx *gin.Context) {
  var user models.User
  err := ctx.ShouldBindJSON(&user)

  if err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{
      "message": "Could not parse request data.",
    })
    return
  }

  err = user.AuthenticateUser()

  if err != nil {
    ctx.JSON(http.StatusUnauthorized, gin.H{
      "message": "Could not authenticate user",
    })
    return
  }

  jwt, err := utils.GenerateToken(user.Email, user.ID)

  if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{
      "message": "Could not authenticate user",
    })
  }

  ctx.JSON(http.StatusOK, gin.H{
    "message": "Login successful",
    "token": jwt,
  })
}
