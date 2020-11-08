package handlers

import (
	"github.com/gin-gonic/gin"	
	"net/http"
	"log"
)

func (handler GameHandler)NewGameHandler (c *gin.Context) {
	game := handler.Service.Start(10,10,10)

	log.Println("hola")
		
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"game": game,
	})
}