package handlers_test

import (
	"net/http"
    "net/http/httptest"
    "testing"
    "github.com/stretchr/testify/assert"
	"minisweeper/api"
	"minisweeper/http/response"	
	"minisweeper/http/request"	
	"encoding/json"	
	"bytes"
)

func TestGameHandler_CreateGameHandlerWithInvalidRequest(t *testing.T) {
	var errorResponse response.ErrorResponse

	requests := []map[string]interface{}{
		{"cols":-123, "rows": -321, "mines": -13},
		{"cols":"some string", "rows": "other string", "mines": "i am a mine"},
		{"cols": 0, "rows": 0, "mines": 0},
	}

	for _, invalidRequest := range requests {
		router := api.InitRoutes()

		out, _ := json.Marshal(invalidRequest)

		req, _ := http.NewRequest("POST", "/api/v1/game/", bytes.NewBuffer(out))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		_ = json.NewDecoder(resp.Body).Decode(&errorResponse)

		expectedResponse := response.ErrorResponse{
			Code: http.StatusBadRequest,
			Message: "error in request",
		}

		if !assert.Equal(t, expectedResponse, errorResponse) {
			t.Log(invalidRequest)
		}
	}
}

func TestCreateGameHandlerWithValidRequest(t *testing.T) {    
	validRequest := request.CreateGameRequest{Cols: 1, Rows: 1, Mines: 1}

	out, _ := json.Marshal(validRequest)
	router := api.InitRoutes()

	req, _ := http.NewRequest("POST", "/api/v1/game/", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)
}

func TestCreateGameHandlerWithEmptyRequest(t *testing.T) {    
	var errorResponse response.ErrorResponse
	
	router := api.InitRoutes()

	req, _ := http.NewRequest("POST", "/api/v1/game/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	
	expectedResponse := response.ErrorResponse{
		Code: 		http.StatusBadRequest,
		Message:	"error in request",
	}

	_ = json.NewDecoder(resp.Body).Decode(&errorResponse)

	assert.Equal(t, errorResponse, expectedResponse)
}