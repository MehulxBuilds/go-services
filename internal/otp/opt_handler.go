package otp

import (
	"fmt"
	"net/http"

	"github.com/MehulxBuilds/go-services/internal/config"
	"github.com/MehulxBuilds/go-services/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service *Service
	cfg     config.Config
}

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func NewHandler(cfg config.Config, service *Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) sendSMS(c *gin.Context) {
	var payload model.OTPModel

	// 1. Bind JSON and check for errors without crashing the server
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed or invalid JSON payload"})
		return // Safely stops processing this single request
	}

	// 2. Validate structural constraints and return errors gracefully
	if err := validate.Struct(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid required fields", "details": err.Error()})
		return // Safely stops processing this single request
	}

	newData := model.OTPModel{
		PhoneNumber: payload.PhoneNumber,
	}

	_, err := h.service.twilioSendOTP(newData.PhoneNumber)
	if err != nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, jsonResponse{Status: http.StatusAccepted, Message: "success", Data: "OTP sent successfully"})
}

func (h *Handler) verifySMS(c *gin.Context) {
	var payload model.VerifyModel

	// 1. Bind JSON and check for errors without crashing the server
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed or invalid JSON payload"})
		return // Safely stops processing this single request
	}

	// 2. Validate structural constraints and return errors gracefully
	if err := validate.Struct(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid required fields", "details": err.Error()})
		return // Safely stops processing this single request
	}

	newData := model.VerifyModel{
		User: payload.User,
		Code: payload.Code,
	}

	err := h.service.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
	fmt.Println("err: ", err)
	if err != nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, jsonResponse{Status: statusCode, Message: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, jsonResponse{Status: http.StatusAccepted, Message: "success", Data: "OTP verified successfully"})
}
