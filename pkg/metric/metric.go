package metric

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		promHandler := promhttp.Handler()

		promHandler.ServeHTTP(c.Writer, c.Request)
	}
}
