package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TwilioAccSid string
	TwilioAuthToken string
	TwilionServiceSid string
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading Env: %v", err)
	}

	cfg := Config{
		TwilioAccSid: getEnv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken: getEnv("TWILIO_AUTHTOKEN"),
		TwilionServiceSid: getEnv("TWILIO_SERVICES_SID"),
	}

	if cfg.TwilioAccSid == "" {
		return Config{}, fmt.Errorf("Acount Sid is required")
	}

	if cfg.TwilioAuthToken == "" {
		return Config{}, fmt.Errorf("Auth Token is required")
	}

	if cfg.TwilionServiceSid == "" {
		return Config{}, fmt.Errorf("Service Sid is required")
	}

	return cfg, nil
}

func getEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("Env value undefined")
	}

	return value
}