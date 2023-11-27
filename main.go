package main

import (
	"log"
	"net/http"
	"os"

	handler "omnial-simulator/handlers"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {

	if err := readEnv(); err != nil {
		log.Fatalf("failed to read env file: %v", err)
	}

	http.HandleFunc("/", handler.InvokeHandle)

	appPort := os.Getenv("APP_PORT")

	log.Println("Omni-simulator invoked successfully")
	http.ListenAndServe(":"+appPort, nil)

}

// Reading environment variables from .env file
func readEnv() error {
	envFilePath := "./env/all.env"

	if err := godotenv.Load(envFilePath); err != nil {
		return err
	}
	return nil
}
