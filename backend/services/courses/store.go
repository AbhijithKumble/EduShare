package courses

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	// Modified query to fetch courses along with the dept name
	getQ := `
	SELECT 
		c.DeptCode, 
		d.DeptName, 
		c.CourseCode, 
		c.CourseName 
	FROM 
		courses c 
	INNER JOIN 
		dept d ON c.DeptCode = d.DeptCode
	`

	// Get a connection to the database
	conn, err := s.db.Conn(c)
	if err != nil {
		log.Fatal("error getting connection to db")
		return nil, err
	}
	defer conn.Close()

	// Execute the query to fetch the courses with the dept name
	rows, err := conn.QueryContext(c, getQ)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("error fetching courses: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and populate the courses slice
	for rows.Next() {
		var course types.Course
		err := rows.Scan(&course.DeptCode, &course.DeptName, &course.CourseCode, &course.CourseName)
		if err != nil {
			log.Print(err)
			return nil, fmt.Errorf("error scanning course row: %w", err)
		}
		courses = append(courses, course)
	}

	// Check for any errors that occurred during row iteration
	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return courses, nil
}

func (s *Store) CreateCourses(c context.Context, courses types.CoursePayload) error {
	checkQ := `SELECT EXISTS(SELECT 1 FROM courses WHERE DeptCode=$1 AND CourseCode=$2)`

	var exists bool

	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return err
	}
	defer conn.Close()

	err = conn.QueryRowContext(c, checkQ, courses.DeptCode, courses.CourseCode).Scan(&exists)

	if err != nil {
		log.Print(err)
		return fmt.Errorf("Course already exits!")
	}

	insertQ := `INSERT INTO courses(DeptCode, CourseCode, CourseName) VALUES  ($1,$2,$3)`
	_, err = conn.ExecContext(c, insertQ, courses.DeptCode, courses.CourseCode, courses.CourseName)

	if err != nil {
		return err
	}

	return nil
}
