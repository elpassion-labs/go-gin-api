package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"

	"go-gin-api/db"
)

func TestCreateTripRouteValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db.Init()
	defer db.CloseDB()

	router := setupRouter()
	body := bytes.NewBuffer([]byte("{\"start_address\":\"Plac Europejski 2, Warszawa\"}"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/trips", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 422, w.Code)
}
