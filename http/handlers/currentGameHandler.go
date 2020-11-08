package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *GameHandler) CurrentGameHandler (c *gin.Context) {
	current := handler.Service.GetCurrent()

	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"game": current,
	})
}
