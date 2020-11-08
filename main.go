package main

import (
	"minisweeper/api"
	"log"
	"os"
)

func main () {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	
	if err != nil {
        log.Fatal(err)
	}
	
	log.SetOutput(file)

	apiRouter := api.InitRoutes()	

	apiRouter.Run()
}