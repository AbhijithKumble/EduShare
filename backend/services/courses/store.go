package courses

import (
	"context"
	"database/sql"

	"github.com/AbhijithKumble/EduShare/backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCourses(c context.Context) ([]types.Course, error) {
	var courses []types.Course

	return courses, nil
}

//func (s *Store) CreateCourses(c context.Context, courses []types.Course) (error) {
//	checkQ := `SELECT EXISTS(SELECT 1 FROM courses WHERE DeptCode+CourseCode=$1)`
//	var exists bool
//
//	conn, err := s.db.Conn(c)
//
//	if err != nil {
//		log.Fatal("error getting connection to db")
//		return err
//	}
//	defer conn.Close()
//
//	err = conn.QueryRowContext(c, checkQ, user.Email).Scan(&exists)
//
//	if err != nil {
//		log.Fatal("error checking if user exists")
//		return err
//	}
//
//	if exists {
//		return fmt.Errorf("user already exists in database")
//	}
//
//	// adding a user into database
//
//	insertQ := `INSERT INTO users(UserID, Email, Password) VALUES
//              ($1, $2, $3)`
//
//	hashedPassword, err := auth.HashPassword(user.Password)
//
//	if err != nil {
//		return err
//	}
//
//	_, err = conn.ExecContext(c, insertQ, uuid.New(), user.Email,
//		hashedPassword)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//  return nil
//}
//
