package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (handler *GameHandler) CurrentGameHandler (c *gin.Context) {
	defer handler.CatchPanic(c)

	//current := handler.Service.GetCurrent()
	//
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		//"game": current,
	})
}
