package routes

import (
    "github.com/gin-gonic/gin"
    "project/controllers"
)

func UserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
    router.POST("/users", userController.CreateUser)
    // Define other routes similarly
}
