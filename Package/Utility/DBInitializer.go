package Utility

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitiateDBInstance(Conn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", Conn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
