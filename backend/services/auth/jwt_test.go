package auth

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateJWT(t *testing.T) {
  id := uuid.New()
  output, err := CreateJWT([]byte("hello"), id)
  
  if err!= nil {
    t.Errorf("Error creating jwt token %v", output) 
  }

  if output == "" {
    t.Errorf("Expected token cannot be empty") 
  } 
}



