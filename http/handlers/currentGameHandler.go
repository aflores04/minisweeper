package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/response"
	"net/http"
)

func (handler *GameHandler) CurrentGameHandler (c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Code:		http.StatusBadRequest,
				Message:	r.(string),
			})
			return
		}
	}()

	current := handler.Service.GetCurrent()

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"game": current,
	})
}
