package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func EnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
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
