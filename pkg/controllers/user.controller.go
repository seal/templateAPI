package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seal/templateapi/pkg/models"
	"github.com/seal/templateapi/pkg/utils"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	currentUser := ctx.Value("currentUser").(models.User)
	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		FirstName: currentUser.FirstName,
		LastName:  currentUser.LastName,
		Email:     currentUser.Email,
		Plan:      currentUser.Plan,
		Username:  currentUser.Username,
	}
	response, err := json.Marshal(userResponse)
	if err != nil {
		err = fmt.Errorf("%w : Error marshalling user response into json", err)
		utils.Error(err)
		utils.HttpError(err, 500, w)
		return
	}
	fmt.Fprint(w, string(response))

	w.Header().Set("Content-Type", "application/json")
}
