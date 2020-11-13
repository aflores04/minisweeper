package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
)

func (handler *GameHandler) CreateGameHandler (c *gin.Context) {
	var postRequest request.CreateGameRequest

	defer handler.CatchPanic(c)

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"error in request",
		})
		return
	}

	game := handler.Service.Start(postRequest.Rows, postRequest.Cols, postRequest.Mines)

	c.JSON(http.StatusOK, response.CreateGameResponse{
		Code: http.StatusOK,
		Game: game,
	})
}