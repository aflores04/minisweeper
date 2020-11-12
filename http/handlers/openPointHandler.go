package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
)

func (handler *GameHandler) OpenPointHandler(c *gin.Context) {
	var pointResponse response.PointResponse

	defer handler.CatchPanic(c)

	var postRequest request.PointRequest

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"error in request",
		})
		return
	}

	pointResponse = handler.Service.OpenPoint(postRequest.Row, postRequest.Col)
	pointResponse.Code = http.StatusOK

	c.JSON(http.StatusOK, pointResponse)
}