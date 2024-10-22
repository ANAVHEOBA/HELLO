package models

import "time"

type User struct {
    ID        string    `json:"id" bson:"_id"`
    Email     string    `json:"email" bson:"email" validate:"required,email"`
    FirstName string    `json:"first_name" bson:"first_name" validate:"required"`
    LastName  string    `json:"last_name" bson:"last_name" validate:"required"`
    Verified  bool      `json:"verified" bson:"verified" validate:"required"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
    UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type UserSecret struct {
    Password string `json:"password" bson:"password" validate:"required"`
    UserID   string `json:"user_id" bson:"user_id" validate:"required"`
}
