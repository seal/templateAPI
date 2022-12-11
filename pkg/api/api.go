package api

import (
	"fmt"
	"log"
	//"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/seal/templateapi/pkg/database"
	"github.com/seal/templateapi/pkg/routes"

	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/seal/templateapi/pkg/controllers"
	"github.com/seal/templateapi/ui"
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
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/api", func(r chi.Router) {

		r.Mount("/admin", AdminRouter())
	})
	r.HandleFunc("/*", indexHandler)
	// static files
	staticFS, err := fs.Sub(ui.StaticFiles, "dist")
	if err != nil {
		log.Println(err, "err here")
	}
	httpFS := http.FileServer(http.FS(staticFS))
	r.Handle("/assets/*", httpFS)
	// Below is a previous version
	/*r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		rawFile, err := ui.StaticFiles.ReadFile("dist/index.html")
		if err != nil {
			log.Println("err in dist/index", err)
		}
		log.Println("IndexHandler used")
		w.Write(rawFile)
	})*/
	/*
		workDir, _ := os.Getwd()
		filesDir := http.Dir(filepath.Join(workDir, "static"))
		FileServer(r, "/", filesDir)
	*/
	return r
}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		rawFile, err := ui.StaticFiles.ReadFile("dist/favicon.ico")
		if err != nil {
			log.Println(err, "error favicon")
		}
		w.Write(rawFile)
		return
	}

	rawFile, err := ui.StaticFiles.ReadFile("dist/index.html")
	if err != nil {
		log.Println("err in dist/index", err)
	}
	w.Write(rawFile)
}

/*
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
}*/

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AdminOnly)
	r.Get("/", AdminAccounts)
	r.Get("/accounts", AdminAccounts)
	/*
		Methods to add:
		/dashboard -> returns tbd, probably total users, total daily searches, etc
	*/
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
