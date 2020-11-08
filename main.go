package main

import (
	"github.com/gin-gonic/gin"
	"minisweeper/api"
	"log"
	"os"
)

func main () {
	router := gin.Default()

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	
	if err != nil {
        log.Fatal(err)
	}
	
	log.SetOutput(file)

	apiRouter := api.InitRoutes(router)	

	apiRouter.Run()
}