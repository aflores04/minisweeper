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

func TestCreateGameHandlerWithValidRequest(t *testing.T) {    
	validRequest := request.CreateGameRequest{
		Cols: 1,
		Rows: 1,
		Mines: 1,
	}

	out, _ := json.Marshal(validRequest)
	
	router := api.InitRoutes()

	req, _ := http.NewRequest("POST", "/game", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)
}

func TestCreateGameHandlerWithEmptyRequest(t *testing.T) {    
	var errorResponse response.ErrorResponse
	
	router := api.InitRoutes()

	req, _ := http.NewRequest("POST", "/game", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	
	expectedResponse := response.ErrorResponse{
		Code: 		http.StatusBadRequest,
		Message:	"error in request",
	}

	err := json.NewDecoder(resp.Body).Decode(&errorResponse)

	if err != nil {
		t.Log(err)
	}

	assert.Equal(t, errorResponse, expectedResponse)
}