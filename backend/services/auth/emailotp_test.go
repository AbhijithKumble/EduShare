package auth

import (
	"testing"
)

func testEmailClientService(t *testing.T) {
	email := "abhijith18765+mailtrap@gmail.com"
	subject := "Verify with Otp"
	body := "Use this code to reset your password"
	otp := "1234"
  err := SendVerificationOtp(email, subject, body, otp)
  if err!=nil {
    t.Fatal("check mail server", err)
    return 
  }
  t.Log("Test case 1: Mail server passed")
}
