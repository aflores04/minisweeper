package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
)

func (handler *GameHandler) CreateGameHandler (c *gin.Context) {
	var postRequest request.CreateGameRequest

	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Code:		http.StatusBadRequest,
				Message:	r.(string),
			})
			return
		}
	}()

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"error in request",
		})
		return
	}

	game := handler.Service.Start(postRequest.Rows, postRequest.Cols, postRequest.Mines)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"game": game,
	})
}