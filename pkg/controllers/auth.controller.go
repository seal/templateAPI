package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/seal/ds/pkg/models"
	"github.com/seal/ds/pkg/utils"
	"github.com/thanhpk/randstr"

	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// [...] SignUp User
func (ac *AuthController) SignUpUser(w http.ResponseWriter, r *http.Request) {
	var payload *models.SignUpInput
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.Error(err)
		utils.HttpError(err, 500, w)
		return
	}
	if payload.Password != payload.PasswordConfirm {
		err = errors.New("Mismatched password")
		utils.Error(err)
		utils.HttpError(err, 401, w)
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		utils.Error(errors.New("Error hashing password"))
		utils.HttpError(err, 401, w)
		return
	}

	newUser := models.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    strings.ToLower(payload.Email),
		Password: hashedPassword,
		Plan:     "free",
		Verified: false,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		utils.Error(errors.New("User with that email already exists"))
		utils.HttpError(err, 401, w)
		return
	} else if result.Error != nil {
		utils.Error(errors.New("User with that email already exists"))

		utils.HttpError(err, 401, w)
		return
	}

	// Generate Verification Code
	code := randstr.String(20)

	verification_code := utils.Encode(code)

	// Update User in Database
	newUser.VerificationCode = verification_code
	ac.DB.Save(newUser)

	var firstName = newUser.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// ? Send Email
	emailData := utils.EmailData{
		URL:       utils.EnvVariable("ClientOrigin") + "/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(&newUser, &emailData)

	message := "We sent an email with a verification code to " + newUser.Email
	fmt.Fprint(w, `{
    "success":true,
    "message":"`+message+`"
    }`)
	w.Header().Set("Content-type", "application/json")
}

// [...] Verify Email

func (ac *AuthController) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("verificationCode")
	verification_code := utils.Encode(code)

	var updatedUser models.User
	result := ac.DB.First(&updatedUser, "verification_code = ?", verification_code)
	if result.Error != nil {
		err := errors.New("Invalid verification code or user doesn't exist")
		utils.Error(err)
		utils.HttpError(err, 401, w)
		return
	}

	if updatedUser.Verified {
		err := errors.New("User already verified")
		utils.Error(err)
		utils.HttpError(err, 200, w) // I know it's an error, but technically it's a good error ?
		return
	}

	updatedUser.VerificationCode = ""
	updatedUser.Verified = true
	ac.DB.Save(&updatedUser)

	fmt.Fprint(w, `{
    "success":true,
    "message":"Email verified successfully"
    }`)
	w.Header().Set("Content-type", "application/json")
}

// [...] SignIn User
func (ac *AuthController) SignInUser(w http.ResponseWriter, r *http.Request) {
	var payload *models.SignInInput

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.Error(err)
		utils.HttpError(err, 501, w)
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		err := errors.New("Invalid email or password")
		utils.Error(err)
		utils.HttpError(err, 401, w)
		return
	}

	if !user.Verified {
		err := errors.New("Please verify your email")
		utils.Error(err)
		utils.HttpError(err, 401, w)
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		err := errors.New("Invalid email or password")
		utils.Error(err)
		utils.HttpError(err, 401, w)
		return
	}

	TokenSecret := utils.EnvVariable("TokenSecret")
	ExpiresIn, err := strconv.Atoi(utils.EnvVariable("TokenExpiresIn"))
	if err != nil {
		err := errors.New("Error parsing enviroment variable for token expires")
		utils.Error(err)
		utils.HttpError(err, 500, w)
		return
	}
	var timeout time.Duration
	timeout = time.Duration(ExpiresIn) * time.Second
	// Generate Token
	token, err := utils.GenerateToken(timeout, user.ID, TokenSecret)
	if err != nil {
		err = fmt.Errorf("%w : Error generating token", err)
		utils.Error(err)
		utils.HttpError(err, 500, w)
		return
	}
	var timenow time.Time
	timenow = time.Now().Add(timeout)
	cookie := &http.Cookie{Name: "token", Path: "/", Expires: timenow, Value: token, HttpOnly: true}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, `"success":true, "token": "`+token+`"`)
}

// [...] SignOut User
func (ac *AuthController) LogoutUser(w http.ResponseWriter, r *http.Request) {
	var expires time.Time
	expires = time.Now()
	cookie := &http.Cookie{Name: "token", Path: "/", Expires: expires, Value: "-1", HttpOnly: true}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, `{
    "success":true,
    "message":"Successfully logged out"
    }`)
}
