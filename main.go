package main

import (

	//"github.com/valyala/fastjson"
	"log"
	"net/http"

	"github.com/seal/ds/pkg/api"

	"github.com/seal/ds/pkg/database"
	"github.com/seal/ds/pkg/utils"
)

func main() {
	database.Connect("scan:" + utils.EnvVariable("MYSQL_SCAN_PASS") + "@tcp(" + utils.EnvVariable("MYSQL_IP") + "):3306)/ds?parseTime=true")
	database.Migrate()
	r := api.GetRouter()

	/*go func() {
		err := http.ListenAndServeTLS(":443",
			"/etc/letsencrypt/live/dsmonkey.com/fullchain.pem",
			"/etc/letsencrypt/live/dsmonkey.com/privkey.pem",
			r)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()*/ // For now no https
	err := http.ListenAndServe(":80",
		r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
