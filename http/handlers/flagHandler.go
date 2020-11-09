package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
)

func (handler *GameHandler) FlagHandler (c *gin.Context) {
	var flagResponse response.FlagResponse

	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Code:		http.StatusBadRequest,
				Message:	r.(string),
			})
			return
		}
	}()

	var postRequest request.FlagRequest

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"error in request",
		})
		return
	}

	flagResponse = handler.Service.AddRemoveFlag(postRequest.Row, postRequest.Col, postRequest.Flag)
	flagResponse.Code = http.StatusOK

	c.JSON(http.StatusBadRequest, flagResponse)
}