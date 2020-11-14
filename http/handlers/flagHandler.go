package handlers

import (
	"github.com/gin-gonic/gin"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
)

func (handler *PointHandler) FlagHandler (c *gin.Context) {
	defer handler.CatchPanic(c)

	var postRequest request.PointRequest

	if err := c.BindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:		http.StatusBadRequest,
			Message:	"error in request",
		})
		return
	}

	point := handler.Service.AddRemoveFlag(postRequest.ID, postRequest.Flag)

	c.JSON(http.StatusOK, response.PointResponse{
		Code: http.StatusOK,
		Point: point,
	})
}