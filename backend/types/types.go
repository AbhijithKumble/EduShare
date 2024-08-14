package types

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserAcc struct {
	UserID                    uuid.UUID `json:"id"`
	Email                     string    `json:"email"`
	Password                  string    `json:"-"`
	CreatedAt                 time.Time `json:"createdAt"`
	UpdatedAt                 time.Time `json:"updatedAt"`
	IsVerified                bool      `json:"isVerified"`
	IsAdmin                   bool      `json:"isAdmin"`
	VerificationToken         *string    `json:"verificationToken"`
	VerificationTokenExpiry   *time.Time `json:"verificationTokenExpiry"`
	ForgotPasswordToken       *string    `json:"forgotPasswordToken"`
	ForgotPasswordTokenExpiry *time.Time `json:"forgotPasswordTokenExpiry"`
}

type UserInfo struct {
	FirstName   string `json:"firstName"`
	MiddleName  string `json:"middleName"`
	LastName    string `json:"lastName"`
	Dept        string `json:"dept"`
	IsDeptAdmin bool   `json:"isDeptAdmin"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(c context.Context, email string) (UserAcc, error)
	CreateUser(c context.Context, user UserAcc) error
}
