package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/seal/templateapi/pkg/database"
	"github.com/seal/templateapi/pkg/models"
	"github.com/seal/templateapi/pkg/utils"
)

func DeserializeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		cookie, err := r.Cookie("token")

		authorizationHeader := r.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie.String()
		}

		if token == "" {
			err := errors.New("You are not logged in ")

			utils.Error(err)
			utils.HttpError(err, 401, w)
			return
		}

		TokenSecret := utils.EnvVariable("TokenSecret")
		sub, err := utils.ValidateToken(token, TokenSecret)
		if err != nil {
			utils.HttpError(err, 401, w)
			utils.Error(err)
			return
		}

		var user models.User
		result := database.Instance.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			err := errors.New("The user belonging to this token no longer exists")
			utils.Error(err)
			utils.HttpError(err, 401, w)
			return
		}

		ctx := context.WithValue(r.Context(), "currentUser", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
