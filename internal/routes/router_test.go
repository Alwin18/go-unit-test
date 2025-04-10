package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Alwin18/go-unit-test/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetRoutes_HelloWorld(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup handler & route config
	helloWorldHandler := handlers.NewHelloWorld()
	routeConfig := NewRouteConfig(helloWorldHandler)
	router := routeConfig.SetRoutes()

	// Perform test request
	req, _ := http.NewRequest("GET", "/hello-world", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"Hello World!"}`, w.Body.String())
}

func TestSetRoutes_Ping(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup handler & route config
	helloWorldHandler := handlers.NewHelloWorld()
	routeConfig := NewRouteConfig(helloWorldHandler)
	router := routeConfig.SetRoutes()

	// Perform test request
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message":"pong"}`, w.Body.String())
}
