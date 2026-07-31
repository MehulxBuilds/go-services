package main

import (
	"log"
	"net/http"

	"github.com/MehulxBuilds/go-services/internal/config"
	"github.com/MehulxBuilds/go-services/internal/otp"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error")
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": "healthy",
		})
	})

	// Register OTP Routes
	otp.RegisterRoutes(cfg, r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed")
	}
}
