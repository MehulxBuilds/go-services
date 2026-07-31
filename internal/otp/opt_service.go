package otp

import (
	"errors"

	"github.com/twilio/twilio-go"
	"github.com/MehulxBuilds/go-services/internal/config"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type Service struct {
	cfg config.Config
}

func NewService(cfg config.Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) getTwilioClient() *twilio.RestClient {
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: s.cfg.TwilioAccSid,
		Password: s.cfg.TwilioAuthToken,
	})

	return client
}

func (s *Service) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := s.getTwilioClient().VerifyV2.CreateVerification(s.cfg.TwilionServiceSid, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (s *Service) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := s.getTwilioClient().VerifyV2.CreateVerificationCheck(s.cfg.TwilionServiceSid, params)
	if err != nil {
		return err
	}

	// BREAKING CHANGE IN THE VERIFY API
	// https://www.twilio.com/docs/verify/quickstarts/verify-totp-change-in-api-response-when-authpayload-is-incorrect
	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}