package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/seal/ds/pkg/controllers"
	"github.com/seal/ds/pkg/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(router chi.Router) {
	/*
		   This "leaks" the middleware, which means the fileserver tries to use user authentication for say, a favicon.ico request.
		   	router.Use(middleware.DeserializeUser)
		   	router.Get("/me", uc.userController.GetMe)
	Issue has been opened on github, below works so no need to change it */

	userGroup := router.Group(nil)
	userGroup.Use(middleware.DeserializeUser)
	userGroup.Get("/user", uc.userController.GetMe)
}
