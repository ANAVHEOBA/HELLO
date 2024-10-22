package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project/models"
    "project/repositories"
    "project/services"
)

type UserController struct {
    repo       *repositories.UserRepository
    jwtService *services.JwtService
}

func NewUserController(repo *repositories.UserRepository, jwtService *services.JwtService) *UserController {
    return &UserController{repo, jwtService}
}

func (uc *UserController) CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := uc.repo.CreateUser(&user, user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

// Additional controller methods like AuthenticateUser, UpdateUser, etc., can be added similarly.
