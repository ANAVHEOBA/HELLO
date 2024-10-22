package repositories

import (
    "context"
    "errors"
    "github.com/jinzhu/gorm"
    "project/models"
    "project/services"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User, password string) error {
    hashedPasswordService := services.NewBcryptService()
    hashedPassword, err := hashedPasswordService.HashPassword(password)
    if err != nil {
        return err
    }

    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    if err := r.db.Create(user).Error; err != nil {
        return err
    }

    secret := &models.UserSecret{Password: hashedPassword, UserID: user.ID}
    return r.db.Create(secret).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, &services.NotFoundError{Message: "User does not exist"}
    }
    return &user, nil
}

// Additional methods like UpdateUser, FindUserById, etc., can be added similarly.
