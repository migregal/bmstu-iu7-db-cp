package statmiddleware

import (
	"neural_storage/pkg/stat"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var responseTimeHistogram stat.HistogramVec

func init() {
	responseTimeHistogram = stat.NewHistogramVec(
		"v1",
		"http_server_request_duration_seconds",
		"Histogram of response time for handler in seconds",
		[]string{"route", "method", "status_code"})
}

func MeasureResponseDuration() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		path := c.FullPath()
		method := c.Request.Method
		status := strconv.Itoa(c.Writer.Status())
		duration := time.Since(start)
		responseTimeHistogram.
			WithLabelValues(path, method, status).
			Observe(float64(duration.Milliseconds()))
	}
}
