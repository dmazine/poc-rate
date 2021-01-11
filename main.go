package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func main() {
	const rate = 5
	const burst = 10

	r := gin.New()
	r.Use(RateLimit(rate, burst))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func RateLimit(r int, b int) gin.HandlerFunc {

	limiter := rate.NewLimiter(rate.Limit(r), b)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}
