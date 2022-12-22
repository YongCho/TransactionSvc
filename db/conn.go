package db

import (
	"database/sql"
	"fmt"

	"pismo.io/util"
)

// GetConnStr returns the DB connection string.
func GetConnStr() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		util.Env.GetDBHost(),
		util.Env.GetDBPort(),
		util.Env.GetDBUser(),
		util.Env.GetDBPassword(),
		util.Env.GetDBName(),
		util.Env.GetDBSSLMode(),
	)
}

// GetDBHandle opens and returns a sql.DB object.
func GetDBHandle() (*sql.DB, error) {
	connStr := GetConnStr()
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
