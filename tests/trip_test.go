package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"

	"go-gin-api/db"
	"go-gin-api/routes"
)

func TestCreateTripRouteValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db.Init()
	defer db.CloseDB()

	router := routes.SetupRouter()
	body := bytes.NewBuffer([]byte("{\"start_address\":\"\"}"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/trips", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
	assert.Contains(t, w.Body.String(), "start address must be present")
}
