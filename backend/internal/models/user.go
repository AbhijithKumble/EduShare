package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID                    uuid.UUID `json:"userid"`
	Username                  string    `json:"username"`
	Email                     string    `json:"email"`
	Password                  string    `json:"-"` 
	Dept                      string    `json:"dept"`
	IsDeptAdmin               bool      `json:"isdeptadmin"`
	IsVerified                bool      `json:"isverified"`
	VerificationToken         string    `json:"verificationtoken"`
	VerificationTokenExpiry   time.Time `json:"verificationtokenexpiry"`
	ForgotPasswordToken       string    `json:"forgotpasswordtoken"`
	ForgotPasswordTokenExpiry time.Time `json:"forgotpasswordtokenexpiry"`
}
