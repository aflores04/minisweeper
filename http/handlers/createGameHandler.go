package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
)

func (handler *GameHandler) CreateGameHandler (c *gin.Context) {
	game := handler.Service.Start(10,10,10)

	var postRequest request.CreateGameRequest

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(400, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"error in request",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"game": game,
	})
}