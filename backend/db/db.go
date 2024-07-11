package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBConnPool struct {
	DB *sql.DB
}

func ConnectDb(dbURL string) (*DBConnPool, error) {
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
        conn.Close()
		return nil, err
	}

	return &DBConnPool{
		DB: conn,
	}, nil
}

func (db *DBConnPool) Close() error{
    if db.DB != nil {
        return db.DB.Close()
    }

    return nil;
}
