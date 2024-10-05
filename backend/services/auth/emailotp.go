package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AbhijithKumble/EduShare/backend/configs"
	"github.com/AbhijithKumble/EduShare/backend/types"
	//"github.com/AbhijithKumble/EduShare/backend/configs"
)

func SendVerificationOtp(email string, subject string, body string, otp string) error {
	email = "abhijith18765+mailtrap@gmail.com"
	url := "https://send.api.mailtrap.io/api/send"
	method := "POST"

	payload := types.EmailPayload{
		Subject:  subject,
		Text:     fmt.Sprintf("%s, Your OTP code is %s", body, otp),
		Category: "Integration Test",
	}
	payload.From.Email = "hello@demomailtrap.com"
	payload.From.Name = "Mailtrap Test"
	payload.To = []struct {
		Email string `json:"email"`
	}{
		{Email: email},
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))

	if err != nil {
		log.Println("Error creating HTTP request:", err)
		return err
	}

	req.Header.Add("Authorization", "Bearer "+string(configs.Envs.EMAIL_CLIENT_AUTH_TOKEN))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
		return err
	}

	defer res.Body.Close()

  _, err = io.ReadAll(res.Body)

	if err != nil {
		log.Println("Error reading response body: ", err)
		return err
	}
  
  if res.StatusCode != 200 {
    return fmt.Errorf("error sending mail")
  }

	return nil
}
