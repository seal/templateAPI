package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID               int    `gorm:"primary_key"`
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	Username         string `json:"username" gorm:"unique"`
	Email            string `json:"email" gorm:"unique"`
	Password         string `json:"password"`
	Plan             string `json:"plan"`
	VerificationCode string
	Verified         bool `gorm:"not null"`
}

type UserPut struct {
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	Username    string `json:"username,omitempty" gorm:"unique"`
	Email       string `json:"email,omitempty" gorm:"unique"`
	NewPasscode string `json:"newpassword,omitempty"`
	OldPasscode string `json:"oldpassword,omitempty"`
}

type SignUpInput struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type UserResponseToken struct {
	User struct {
		ID        int    `json:"ID"`
		FirstName string `json:"firstname" binding:"required"`
		LastName  string `json:"lastname" binding:"required"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		Plan      string `json:"plan"`
	} `json:"user"`
	Token string `json:"token"`
}
type UserResponse struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email,omitempty"`
	Plan      string `json:"role,omitempty"`
	Username  string `json:"username,omitempty"`
}
