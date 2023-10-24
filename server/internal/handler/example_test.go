package handler

import (
	"api-gateway-template/app"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	// mock database GetExampleData method
	resp := &app.Data{IsFake: true, Date: time.Now()}
	fakeAPI.Fetcher.On("GetExampleData", mock.Anything).Return(resp, nil)

	req := httptest.NewRequest(http.MethodGet, "/example", nil)
	w := httptest.NewRecorder()
	muxRouter.ServeHTTP(w, req)

	var got *app.Data
	_ = json.NewDecoder(w.Body).Decode(&got)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, resp.IsFake, got.IsFake)

}
