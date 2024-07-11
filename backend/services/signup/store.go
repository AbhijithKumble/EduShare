package signup

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/AbhijithKumble/EduShare/backend/services/auth"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(c context.Context, email string, password string) error {
	checkQ := `SELECT EXISTS(SELECT 1 FROM users WHERE Email=$1)`
	var exists bool

	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return err
	}

	defer conn.Close()

	err = conn.QueryRowContext(c, checkQ, email).Scan(&exists)

	if err != nil {
		log.Fatal("error checking if user exists")
		return err
	}

	if exists {
		return fmt.Errorf("user already exists in database")
	}

	// adding a user into database

	insertQ := `INSERT INTO users(UserID, CreatedAt, UpdatedAt, Email, Password)
                VALUES ($1, $2, $3, $4, $5)`

    hashedPassword, err := auth.HashPassword(password)
    if err!=nil {
        return err
    }

    _, err = conn.ExecContext(c, insertQ, uuid.New(), time.Now().UTC(), time.Now().UTC(), email, hashedPassword)

    if err!=nil {
        return err
    }

	return nil
}
