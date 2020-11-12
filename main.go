package main

import (
	"github.com/joho/godotenv"
	"minisweeper/api"
	"log"
	"os"
)

func main () {
	err := godotenv.Load("./.env.local")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	
	if err != nil {
        log.Fatal(err)
	}
	
	log.SetOutput(file)

	apiRouter := api.InitRoutes()	

	apiRouter.Run()
}