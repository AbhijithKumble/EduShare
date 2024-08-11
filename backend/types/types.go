package types

import (
	"context"
	"time"
)

type User struct {
	UserID                    int       `json:"id"`
	FirstName                 string    `json:"firstName"`
	MiddleName                string    `json:"middleName"`
	LastName                  string    `json:"lastName"`
	Dept                      string    `json:"dept"`
	Email                     string    `json:"email"`
	Password                  string    `json:"-"`
	CreatedAt                 time.Time `json:"createdAt"`
	UpdatedAt                 time.Time `json:"updatedAt"`
	IsDeptAdmin               bool      `json:"isDeptAdmin"`
	IsVerified                bool      `json:"isVerified"`
	VerificationToken         string    `json:"verificationToken"`
	VerificationTokenExpiry   time.Time `json:"verificationTokenExpiry"`
	ForgotPasswordToken       string    `json:"forgotPasswordToken"`
	ForgotPasswordTokenExpiry time.Time `json:"forgotPasswordTokenExpiry"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateUser(c context.Context, user User) error
}

