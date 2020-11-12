package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"minisweeper/http/request"
	"minisweeper/http/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenPointHandler_SuccessResponse(t *testing.T)  {
	validRequest := request.PointRequest{
		Col: 1,
		Row: 1,
	}

	var successResponse response.PointResponse

	out, _ := json.Marshal(validRequest)
	router := SetUpRouter()
	req, _ := http.NewRequest("PUT", "/api/v1/game/point/open", bytes.NewBuffer(out))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	_ = json.NewDecoder(resp.Body).Decode(&successResponse)

	assert.Equal(t, true, successResponse.Open)
	assert.Equal(t, http.StatusOK, resp.Code)
}