package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
	pointRepository := repositories.NewPointRepository(connection)

	gameService 	:= services.NewGameService(gameRepository)
	pointService 	:= services.NewPointService(pointRepository)

	gameHandler 	:= handlers.NewGameHandler(gameService)
	pointHandler 	:= handlers.NewPointHandler(pointService)

	// init a game
	gameRepository.Create(2,2,1)

	v1 := router.Group("/api/v1")

	game := v1.Group("game")
	game.POST("/", gameHandler.CreateGameHandler)
	game.GET("/", gameHandler.CurrentGameHandler)

	point := game.Group("point")
	point.PUT("flag", pointHandler.FlagHandler)
	//point.PUT("open", gameHandler.OpenPointHandler)

	return router
}

func TestGameHandler_FlagHandlerWithBadRequest(t *testing.T) {
	var errorResponse response.ErrorResponse

	requests := []map[string]interface{}{
		{"id": 99999, "col": 0, "row": 0, "flag": 12},
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
	var (
		successResponse response.PointResponse
		out []byte
		req *http.Request
		resp *httptest.ResponseRecorder
	)

	gameRepository := repositories.NewGameRepository(database.NewConnectionTest())

	game := gameRepository.Create(1,1,1)

	router := SetUpRouter()

	trueRequest 	:= request.PointRequest{ID: game.Points[0].ID, Flag: true}
	falseRequest 	:= request.PointRequest{ID: game.Points[0].ID, Flag: false}

	out, _ = json.Marshal(trueRequest)
	req, _ = http.NewRequest("PUT", "/api/v1/game/point/flag", bytes.NewBuffer(out))
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	_ = json.NewDecoder(resp.Body).Decode(&successResponse)

	assert.Equal(t, true, successResponse.Point.Flag)
	assert.Equal(t, http.StatusOK, resp.Code)

	out, _ = json.Marshal(falseRequest)
	req, _ = http.NewRequest("PUT", "/api/v1/game/point/flag", bytes.NewBuffer(out))
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	_ = json.NewDecoder(resp.Body).Decode(&successResponse)

	assert.Equal(t, false, successResponse.Point.Flag)
	assert.Equal(t, http.StatusOK, resp.Code)
}
