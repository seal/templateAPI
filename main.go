package main

import (
	"log"
	"net/http"

	"github.com/seal/templateapi/pkg/api"

	"github.com/seal/templateapi/pkg/database"
	"github.com/seal/templateapi/pkg/utils"
)

func main() {
	database.Connect(utils.EnvVariable("MYSQL_USERNAME") + ":" + utils.EnvVariable("MYSQL_PASSWORD") + "@tcp(" + utils.EnvVariable("MYSQL_HOST") + ":3306)/" + utils.EnvVariable("MYSQL_DATABASENAME") + "?parseTime=true")
	database.Migrate()
	r := api.GetRouter()
	/*chi.Walk(r, func(method, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		log.Printf("%-10s | %s\n", method, route)
		return nil
	})*/
	// Use above to print all routes on startup
	/*
		go func() {

			err := http.ListenAndServeTLS(":443",
				"/etc/letsencrypt/live/" + utils.EnvVariable("DOMAIN_HOST") + "/fullchain.pem",
				"/etc/letsencrypt/live/" + utils.EnvVariable("DOMAIN_HOST") + "/privkey.pem",
				r)
			if err != nil {
				log.Fatal("ListenAndServe: ", err)
			}
		}()*/
	err := http.ListenAndServe(":80",
		r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
