package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/AbhijithKumble/EduShare/backend/services/auth"
	"github.com/AbhijithKumble/EduShare/backend/types"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(c context.Context, user types.UserAcc) (*uuid.UUID, error) {
    checkQVerified := `SELECT EXISTS (SELECT UserID FROM users WHERE Email=$1 AND IsVerified = true)`
    checkQ := `SELECT EXISTS(SELECT 1 FROM users WHERE Email=$1)`

    var exists bool

    conn, err := s.db.Conn(c)
    if err != nil {
        log.Fatal("error getting connection to db")
        return nil, err
    }
    defer conn.Close()

    // Check if a verified user exists
    err = conn.QueryRowContext(c, checkQVerified, user.Email).Scan(&exists)

    if err != nil {
        log.Fatal("error checking if user exists")
        return nil, err
    }
    if exists == true {
        return nil, fmt.Errorf("user already exists in the database")
    }

    // Check if the user exists
    err = conn.QueryRowContext(c, checkQ, user.Email).Scan(&exists)

    if err != nil {
        log.Fatal("error checking if user exists")
        return nil, err
    }

    // Prepare to create user or update password if already present
    insertQ := `INSERT INTO users (UserID, Email, Password, VerificationToken, VerificationTokenExpiry)
                 VALUES ($1, $2, $3, $4, $5)
                 ON CONFLICT (Email) 
                 DO UPDATE SET 
                     Password = EXCLUDED.Password,
                     VerificationToken = EXCLUDED.VerificationToken,
                     VerificationTokenExpiry = EXCLUDED.VerificationTokenExpiry
                 RETURNING UserID`  // Return the UserID on conflict as well

    hashedPassword, err := auth.HashPassword(user.Password)
    if err != nil {
        return nil, err
    }

    // Generate verification token and expiration time
    verifyCode := math.Floor(100000 + rand.Float64()*900000)
    verificationToken := strconv.FormatFloat(verifyCode, 'f', 0, 64)
    expirationTime := time.Now().Add(30 * time.Minute)

    var userID uuid.UUID  // Variable to hold the returned UserID

    // Execute the insert query and get the UserID
    err = conn.QueryRowContext(c, insertQ, uuid.New(), user.Email, hashedPassword, verificationToken, expirationTime).Scan(&userID)
    if err != nil {
        return nil, err
    }

    otpSubject := "OTP to verify your EduShare account"
    mailBody := "This is your OTP"

    err = auth.SendVerificationOtp(user.Email, otpSubject, mailBody, verificationToken)
    if err != nil {
        log.Printf("error in sending OTP")
        return nil, err
    }

    return &userID, nil  // Return the pointer to the collected UserID
}

func (s *Store) VerifyOtp(c context.Context, email string, password string, otp string) (int, error) {
	user, err, _ := s.GetUserByEmail(c, email)

	if err != nil {
		return 500, err
	}

	if !auth.ComparePassword(user.Password,password) {
		return 400, fmt.Errorf("invalid password")
	}

	if user.VerificationTokenExpiry == nil {
		return 404, fmt.Errorf("invalid token")
	}

	if (user.VerificationTokenExpiry).Before(time.Now()) {
		return 404, fmt.Errorf("invalid token")
	}

	if user.VerificationToken == nil {
		return 400, fmt.Errorf("token not found")
	}

	if *(user.VerificationToken) != otp {
		return 404, fmt.Errorf("invalid token")
	}

  	// If all checks pass, mark the user as verified
	updateQ := `UPDATE users SET IsVerified = true WHERE UserID = $1`
	_, err = s.db.ExecContext(c, updateQ, user.UserID)

	if err != nil {
		return 500, fmt.Errorf("error updating verification status")
	}

	return 200, nil
}

func (s *Store) GetUserByEmail(c context.Context, email string) (types.UserAcc, error, int) {

	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return types.UserAcc{}, err, 500
	}

	defer conn.Close()

	query := `SELECT * FROM users WHERE Email =$1`

	row := conn.QueryRowContext(c, query, string(email))

	var user types.UserAcc

	err = row.Scan(&user.UserID, &user.Email, &user.Password, &user.CreatedAt,
		&user.UpdatedAt, &user.IsVerified, &user.IsAdmin,
		&user.VerificationToken, &user.VerificationTokenExpiry,
		&user.ForgotPasswordToken, &user.ForgotPasswordTokenExpiry)

	switch {
	case err == sql.ErrNoRows:
		return types.UserAcc{}, fmt.Errorf("user not found"), 404
	case err != nil:
		return types.UserAcc{}, err, 500
	default:
		return user, nil, 200
	}
}

//-- duplicate thing

//func (s *Store) verifyUser(c context.Context, userInfo types.LoginUserPayload, otp string) (error, int) {
//
//	conn, err := s.db.Conn(c)
//
//	if err != nil {
//		log.Fatal("error getting connection to db")
//		return err, 500
//	}
//
//	defer conn.Close()
//
//	query := `SELECT * FROM users WHERE Email =$1`
//
//	row := conn.QueryRowContext(c, query, string(userInfo.Email))
//
//	var user types.UserAcc
//
//	err = row.Scan(&user.UserID, &user.Email, &user.Password, &user.CreatedAt,
//		&user.UpdatedAt, &user.IsVerified, &user.IsAdmin,
//		&user.VerificationToken, &user.VerificationTokenExpiry,
//		&user.ForgotPasswordToken, &user.ForgotPasswordTokenExpiry)
//
//	if err != nil {
//		return fmt.Errorf("error getting info of user"), 500
//	}
//
//	if user.VerificationTokenExpiry.Before(time.Now()) {
//		return fmt.Errorf("invalid token"), 400
//	}
//
//	if otp != *user.VerificationToken {
//		return fmt.Errorf("invalid token"), 400
//	}
//
//	return nil, 200
//}

//func (s *Store) UpdateUserInfo(c context.Context, userInfo types.UserInfo) (error) {
//
//	conn, err := s.db.Conn(c)
//
//	if err != nil {
//		log.Fatal("error getting connection to db")
//		return  err
//	}
//
//	defer conn.Close()
//
//	query := `SELECT * FROM users WHERE Email =$1`
//
//	row := conn.QueryRowContext(c, query, string(email))
//
//  var user types.UserAcc
//
//	err = row.Scan(&user.UserID, &user.Email, &user.Password, &user.CreatedAt,
//        &user.UpdatedAt, &user.IsVerified, &user.IsAdmin,
//        &user.VerificationToken, &user.VerificationTokenExpiry,
//        &user.ForgotPasswordToken, &user.ForgotPasswordTokenExpiry)
//
//	switch {
//
//	case err == sql.ErrNoRows:
//		return types.UserAcc{}, fmt.Errorf("user not found")
//	case err != nil:
//		return types.UserAcc{}, err
//	default:
//		return user, nil
//	}
//
//}
