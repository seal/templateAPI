package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/seal/ds/pkg/controllers"
	"github.com/seal/ds/pkg/middleware"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRouter(router chi.Router) {
	router.Post("/register", rc.authController.SignUpUser)
	router.Post("/login", rc.authController.SignInUser)
	router.Get("/verifyemail/:verificationCode", rc.authController.VerifyEmail)
	//r.Get("/logout", rc.authController.LogoutUser.ServeHTTP)

	logoutGroup := router.Group(nil)
	logoutGroup.Use(middleware.DeserializeUser)
	logoutGroup.Get("/logout", rc.authController.LogoutUser)
	/*
		router.Group(func(r chi.Router) {
			r.Use(middleware.DeserializeUser)
			r.Get("/logout", rc.authController.LogoutUser)
		})*/
}
