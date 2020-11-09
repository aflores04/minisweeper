package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"minisweeper/api"
	"minisweeper/http/handlers"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"minisweeper/repositories"
	"minisweeper/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpRouter() *gin.Engine  {
	router := gin.Default()
	gameRepository 	:= repositories.NewGameRepository()
	gameService 	:= services.NewGameService(gameRepository)
	gameHandler 	:= handlers.NewGameHandler(gameService)

	// init a game
	gameRepository.CreateGame(2,2,0)

	router.POST("/game/flag", gameHandler.FlagHandler)

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
		router := setUpRouter()

		out, _ := json.Marshal(invalidRequest)
		req, _ := http.NewRequest("POST", "/game/flag", bytes.NewBuffer(out))
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
	var successResponse response.FlagResponse
	validRequest := request.FlagRequest{Col: 1, Row: 1, Flag: true}

	router := setUpRouter()
	out, _ := json.Marshal(validRequest)
	req, _ := http.NewRequest("POST", "/game/flag", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	expectedResponse := response.FlagResponse{Code: http.StatusOK, Row: 1, Col: 1, Mine: false, Flag: true}

	_ = json.NewDecoder(resp.Body).Decode(&successResponse)

	assert.Equal(t, expectedResponse, successResponse)
}

func TestGameHandler_FlagHandlerWithNoGame(t *testing.T) {
	var errorResponse response.ErrorResponse

	validRequest := request.FlagRequest{Col: 1, Row: 1, Flag: true,}

	out, _ := json.Marshal(validRequest)
	router := api.InitRoutes()
	req, _ := http.NewRequest("POST", "/game/flag", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	expectedResponse := response.ErrorResponse{Code: http.StatusBadRequest, Message: "there is no game running",}

	_ = json.NewDecoder(resp.Body).Decode(&errorResponse)

	assert.Equal(t, expectedResponse, errorResponse)
}
