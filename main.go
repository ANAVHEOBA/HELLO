package main

import (
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "project/controllers"
    "project/repositories"
    "project/routes"
    "project/services"
)

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect to the database")
    }
    defer db.Close()

    db.AutoMigrate(&models.User{}, &models.UserSecret{})

    jwtService := services.NewJwtService("your_secret_key")
    userRepository := repositories.NewUserRepository(db)
    userController := controllers.NewUserController(userRepository, jwtService)

    router := gin.Default()
    api := router.Group("/api")
    routes.UserRoutes(api, userController)

    router.Run(":8080")
}
