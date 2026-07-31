package otp

import (
	"github.com/MehulxBuilds/go-services/internal/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(cfg config.Config, r *gin.Engine) {
	// create repo and handler once at startup
	s := NewService(cfg)
	h := NewHandler(cfg, s)

	otpGroup := r.Group("/otp")
	{
		otpGroup.POST("/", h.sendSMS)
		otpGroup.POST("/verifyOtp", h.verifySMS)
	}
}
