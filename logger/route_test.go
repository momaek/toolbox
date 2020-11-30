package logger

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinLogger(t *testing.T) {
	engine := gin.Default()
	engine.Use(GinLoggerMiddleware())
	engine.GET("/", func(c *gin.Context) {
		log := GinRequestLogger(c)
		log.Info("hello")
	})

	go engine.Run(":9090")

	http.Get("http://localhost:9090/")
}
