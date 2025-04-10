package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Alwin18/go-unit-test/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	helloWorldHandler := handlers.NewHelloWorld()
	routeConfig := NewRouteConfig(helloWorldHandler)
	router = routeConfig.SetRoutes()

	os.Exit(m.Run())
}

func TestSetRoutes_HelloWorld(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello-world", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Hello World!"}`, w.Body.String())
}

func TestSetRoutes_Ping(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
}
