package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               int    `gorm:"primary_key"`
	Name             string `json:"name"`
	Username         string `json:"username" gorm:"unique"`
	Email            string `json:"email" gorm:"unique"`
	Password         string `json:"password"`
	Plan             string `json:"plan"`
	VerificationCode string
	Verified         bool `gorm:"not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Plan     string `json:"role,omitempty"`
	Username string `json:"username,omitempty"`
}
