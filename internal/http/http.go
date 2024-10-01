package http

import (
	"kompack-go-api/internal/factory"
	"kompack-go-api/internal/middleware"
	"kompack-go-api/pkg/metric"
	"kompack-go-api/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHttp(g *gin.Engine, f *factory.Factory) {
	g.Use(middleware.CORS())
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	// index route
	g.GET("/", index)

	// metrics route
	g.GET("/kompack-api-metrics", metric.PrometheusHandler())

	// v1 := g.Group("/api/v1")
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{
		Name:    "kompack-go-api",
		Version: util.LoadVersion(),
	})
}
