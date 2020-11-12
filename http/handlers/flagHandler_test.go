package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"minisweeper/api"
	"minisweeper/database"
	"minisweeper/http/handlers"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"minisweeper/repositories"
	"minisweeper/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine  {
	router := gin.Default()
	connection := database.NewConnectionTest()
	gameRepository 	:= repositories.NewGameRepository(connection)
	gameService 	:= services.NewGameService(gameRepository)
	gameHandler 	:= handlers.NewGameHandler(gameService)

	// init a game
	gameRepository.Create(2,2,1)

	v1 := router.Group("/api/v1")

	game := v1.Group("game")
	game.POST("/", gameHandler.CreateGameHandler)
	game.GET("/", gameHandler.CurrentGameHandler)

	point := game.Group("point")
	point.PUT("flag", gameHandler.FlagHandler)
	point.PUT("open", gameHandler.OpenPointHandler)

	return router
}

func TestGameHandler_FlagHandlerWithBadRequest(t *testing.T) {
	var errorResponse response.ErrorResponse

	requests := []map[string]interface{}{
		{"col": 0, "row": 0, "flag": 12},
		{"col": "some string", "row": "other", "flag": "hello world"},
		{"col": -123, "row": 123, "flag": true},
	}

	for _, invalidRequest := range requests {
		router := SetUpRouter()

		out, _ := json.Marshal(invalidRequest)
		req, _ := http.NewRequest("PUT", "/api/v1/game/point/flag", bytes.NewBuffer(out))
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		_ = json.NewDecoder(resp.Body).Decode(&errorResponse)

		expectedResponse := response.ErrorResponse{Code: 400, Message: "error in request"}

		if ! assert.Equal(t, expectedResponse, errorResponse) {
			t.Log(invalidRequest)
		}
	}

}

func TestGameHandler_FlagHandlerSuccessResponse(t *testing.T) {
	var successResponse response.PointResponse
	validRequest := request.PointRequest{Col: 1, Row: 1, Flag: true}

	router := SetUpRouter()
	out, _ := json.Marshal(validRequest)
	req, _ := http.NewRequest("PUT", "/api/v1/game/point/flag", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	expectedResponse := response.PointResponse{Code: http.StatusOK, Row: 1, Col: 1, Value: 1, Mine: false, Flag: true}

	_ = json.NewDecoder(resp.Body).Decode(&successResponse)

	assert.Equal(t, expectedResponse, successResponse)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGameHandler_FlagHandlerWithNoGame(t *testing.T) {
	var errorResponse response.ErrorResponse

	validRequest := request.PointRequest{Col: 1, Row: 1, Flag: true,}

	out, _ := json.Marshal(validRequest)
	router := api.InitRoutes()
	req, _ := http.NewRequest("PUT", "/api/v1/game/point/flag", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	expectedResponse := response.ErrorResponse{Code: http.StatusBadRequest, Message: "there is no game running",}

	_ = json.NewDecoder(resp.Body).Decode(&errorResponse)

	assert.Equal(t, expectedResponse, errorResponse)
}
