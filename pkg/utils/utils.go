package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/seal/templateapi/pkg/types"
)

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
func CountryLoopup(r *http.Request) string {
	ip := strings.Split(ReadUserIP(r), ":")
	resp, err := http.Get("https://api.iplocation.net?ip=" + ip[0])
	if err != nil {
		Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Error(err)
	}
	var IPResponse types.IpLoopup
	json.Unmarshal(body, &IPResponse)
	return IPResponse.CountryCode2
}
func EnvVariable(key string) string {
	/*
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		} */
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)

	}
	return os.Getenv(key)
}
func HttpError(err error, StatusCode int, w http.ResponseWriter) {
	ErrorStruct := &ErrorResponse{
		Success: false,
		Message: err.Error(),
	}
	response, err := json.Marshal(ErrorStruct)
	if err != nil {
		Error(err) // Yeah not very good practice
	}

	w.WriteHeader(StatusCode)
	w.Write(response)
}

type ErrorResponse struct { // Add to models package later
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Error(err error) {
	log.Println(err)
	SendTelegramMessage(err.Error())
}

func SendTelegramMessage(msg string) {

	url := "https://api.telegram.org/bot" + EnvVariable("TELEGRAM_BOT_TOKEN") + "/sendMessage?chat_id=" + EnvVariable("TELEGRAM_CHAT_ID") + "&text=" + msg
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
