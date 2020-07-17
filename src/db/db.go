package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// fixQueryParams turns a query containing '?' characters into '$1, $2, $3, etc.'
func fixQueryParams(sql string) string {
	for paramCount := 1; strings.Contains(sql, "?"); paramCount++ {
		sql = strings.Replace(sql, "?", fmt.Sprintf("$%v", paramCount), 1)
	}
	return sql
}

// DBManager represents a database manager
type DBManager struct {
	dbURL string
	pool  *pgxpool.Pool
}

// NewDBManager creates a new database manager object
func NewDBManager(dbURL string) (*DBManager, error) {
	dbpool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	return &DBManager{
		dbURL: dbURL,
		pool:  dbpool,
	}, nil
}

// Execute executes a sql query
func (dbm *DBManager) Execute(sql string, params ...interface{}) error {
	_, err := dbm.pool.Exec(context.Background(), fixQueryParams(sql), params...)
	return err
}

// QueryRow executes a sql query that returns a row
func (dbm *DBManager) QueryRow(sql string, params ...interface{}) pgx.Row {
	fmt.Println(fixQueryParams(sql))
	return dbm.pool.QueryRow(context.Background(), fixQueryParams(sql), params...)
}

// QueryRows executes a sql query that returns multiple rows
func (dbm *DBManager) QueryRows(sql string, params ...interface{}) (pgx.Rows, error) {
	return dbm.pool.Query(context.Background(), fixQueryParams(sql), params...)
}

// Close closes the database pool
func (dbm *DBManager) Close() {
	dbm.pool.Close()
}
