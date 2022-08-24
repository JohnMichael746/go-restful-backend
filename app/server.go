package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AlxPatidar/go-restful-api/router"
	"github.com/joho/godotenv"
)

func init() {
	error := godotenv.Load("config/.env")
	if error != nil {
		log.Fatal("Error loading .env file.")
		fmt.Printf("Error loading .env file %s", error.Error())
		return
	}
}

func Start() {
	// get port from env file
	port := os.Getenv("PORT")
	fmt.Println("Go server is started on", port)
	router := router.Router()
	// run server on env port
	http.ListenAndServe(":"+port, router)
}
