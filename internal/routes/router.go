package routes

import (
	"github.com/Alwin18/go-unit-test/internal/handlers"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	helloWorldHandler *handlers.HelloWorld
}

func NewRouteConfig(helloWorldHandler *handlers.HelloWorld) *RouteConfig {
	return &RouteConfig{
		helloWorldHandler: helloWorldHandler,
	}
}

func (rc *RouteConfig) SetRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello-world", rc.helloWorldHandler.HelloWorld)

	return r
}
