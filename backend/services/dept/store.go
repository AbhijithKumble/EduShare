package dept

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

func (s *Store) GetDepts(c context.Context) ([]types.Dept, error) {
	var depts []types.Dept

	// Define the query to get all departments
	getQ := `SELECT DeptCode, DeptName FROM dept`

	// Get a connection to the database
	conn, err := s.db.Conn(c)
	if err != nil {
		log.Fatal("error getting connection to db")
		return nil, err
	}
	defer conn.Close()

	// Execute the query to get departments
	rows, err := conn.QueryContext(c, getQ)
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("error fetching departments: %w", err)
	}
	defer rows.Close()

	// Iterate over the rows and populate the depts slice
	for rows.Next() {
		var dept types.Dept
		err := rows.Scan(&dept.DeptCode, &dept.DeptName)
		if err != nil {
			log.Print(err)
			return nil, fmt.Errorf("error scanning department row: %w", err)
		}
		depts = append(depts, dept)
	}

	// Check for any errors that occurred during row iteration
	if err := rows.Err(); err != nil {
		log.Print(err)
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return depts, nil
}

func (s *Store) CreateDepts(c context.Context, dept types.DeptPayload) error {
	checkQ := `SELECT EXISTS(SELECT 1 FROM dept WHERE DeptCode=$1)`

	var exists bool

	conn, err := s.db.Conn(c)

	if err != nil {
		log.Fatal("error getting connection to db")
		return err
	}
	defer conn.Close()

	err = conn.QueryRowContext(c, checkQ, dept.DeptCode).Scan(&exists)

	if err != nil {
		log.Print(err)
		return fmt.Errorf("Dept already exits!")
	}

	insertQ := `INSERT INTO dept(DeptCode, DeptName) VALUES  ($1,$2)`
	_, err = conn.ExecContext(c, insertQ, dept.DeptCode, dept.DeptName)

	if err != nil {
		return err
	}

	return nil
}
