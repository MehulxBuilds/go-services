package model

type OTPModel struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
}

type VerifyModel struct {
	User *OTPModel `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
