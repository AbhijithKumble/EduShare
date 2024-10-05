package auth

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AbhijithKumble/EduShare/backend/configs"
	"github.com/AbhijithKumble/EduShare/backend/types"
	"github.com/AbhijithKumble/EduShare/backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateJWT(secret []byte, userID uuid.UUID) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWT_EXPIRATION_IN_SECONDS)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func WithJWT(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    tokenString := utils.GetTokenFromRequest(r)

    token, err := validateJWT(tokenString) 
    
    if err != nil {
      log.Printf("failed to validate token : %v", err)
      permissionDenied(w)
      return
    }

    if !token.Valid {
      log.Print("invalid token")
      permissionDenied(w)
      return 
    } 

    claims := token.Claims.(jwt.MapClaims) 
    str := claims["userID"].(string)
    if str != "" {}
    
    handlerFunc(w, r)

	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Envs.JWT_SECRET), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}
