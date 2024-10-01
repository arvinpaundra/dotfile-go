package http

import (
	"net/http"

	"github.com/arvinpaundra/dotfile-go/internal/factory"
	"github.com/arvinpaundra/dotfile-go/internal/middleware"
	"github.com/arvinpaundra/dotfile-go/pkg/metric"
	"github.com/arvinpaundra/dotfile-go/pkg/util"

	"github.com/gin-gonic/gin"
)

func NewHttp(g *gin.Engine, f *factory.Factory) {
	g.Use(middleware.CORS())
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	// index route
	g.GET("/", index)

	// metrics route
	g.GET("/api-metrics", metric.PrometheusHandler())

	// v1 := g.Group("/api/v1")
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{
		Name:    "dotfile-go",
		Version: util.LoadVersion(),
	})
}
