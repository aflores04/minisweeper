package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/response"
	"minisweeper/services"
	"net/http"
)

type GameHandler struct {
	Service services.IGameService
}

func NewGameHandler(service services.IGameService) *GameHandler {
	return &GameHandler{
		Service: service,
	}
}

func (GameHandler) CatchPanic(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	r.(string),
		})
		return
	}
}



