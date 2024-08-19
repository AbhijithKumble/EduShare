package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func (s *Store) CreateUser(c context.Context, user types.UserAcc) error {
	checkQ := `SELECT EXISTS(SELECT 1 FROM users WHERE Email=$1)`
	var exists bool

	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return err
	}
	defer conn.Close()

	err = conn.QueryRowContext(c, checkQ, user.Email).Scan(&exists)

	if err != nil {
		log.Fatal("error checking if user exists")
		return err
	}

	if exists {
		return fmt.Errorf("user already exists in database")
	}

	// adding a user into database

	insertQ := `INSERT INTO users(UserID, Email, Password) VALUES 
              ($1, $2, $3)`

	hashedPassword, err := auth.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = conn.ExecContext(c, insertQ, uuid.New(), user.Email,
		hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByEmail(c context.Context, email string) (types.UserAcc, error) {

	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return types.UserAcc{}, err
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
		return types.UserAcc{}, fmt.Errorf("user not found")
	case err != nil:
		return types.UserAcc{}, err
	default:
		return user, nil
	}
}

func (s *Store) UpdateUserInfo(c context.Context, userInfo types.UserInfo) (error) {
  
	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return  err
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
		return types.UserAcc{}, fmt.Errorf("user not found")
	case err != nil:
		return types.UserAcc{}, err
	default:
		return user, nil
	}


  return nil
}
