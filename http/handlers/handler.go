package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/response"
	"minisweeper/services"
	"net/http"
)

type IHandler interface {
	CatchPanic(c *gin.Context)
}

type Handler struct {}

type GameHandler struct {
	Handler
	Service services.IGameService
}

type PointHandler struct {
	Handler
	Service services.IPointService
}

func NewGameHandler(service services.IGameService) *GameHandler {
	return &GameHandler{
		Service: service,
	}
}

func NewPointHandler(service services.IPointService) *PointHandler  {
	return &PointHandler{
		Service: service,
	}
}

func (Handler) CatchPanic(c *gin.Context) {
	if r := recover(); r != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	r.(string),
		})
		return
	}
}



