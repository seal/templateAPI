package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/seal/ds/pkg/database"
	"github.com/seal/ds/pkg/routes"

	"github.com/seal/ds/pkg/controllers"

	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func GetRouter() *chi.Mux {
	AuthController = controllers.NewAuthController(database.Instance)
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	UserController = controllers.NewUserController(database.Instance)
	UserRouteController = routes.NewRouteUserController(UserController)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/user", func(r chi.Router) {

		UserRouteController.UserRoute(r)
		AuthRouteController.AuthRouter(r)
	})
	r.Mount("/api/admin", AdminRouter())
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "js/dist"))
	FileServer(r, "/", filesDir)
	return r
}
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Get("/", AdminAccounts)
	r.Get("/accounts", AdminAccounts)
	return r
}

func AdminAccounts(w http.ResponseWriter, r *http.Request) {
	// Example router, for now ignore
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Its ok "))
}

type AdminLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		// e.g.
		// if !isAdmin(r) {
		// 	http.Error(w, http.StatusText(401), 401)
		// 	return
		// }
		IsAdmin, _ := isAdmin(r)
		if !IsAdmin {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func isAdmin(r *http.Request) (bool, AdminLoginResponse) {
	var RemoveLater AdminLoginResponse
	return true, RemoveLater
}
