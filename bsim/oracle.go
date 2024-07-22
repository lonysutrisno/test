package bsim

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	go_ora "github.com/sijms/go-ora/v2"
)

func ConnectDB() *sqlx.DB {
	dsn := go_ora.BuildUrl("10.22.103.79", 1521, "orcl", "SIMASBANKINGPROD", "simasbanking", nil)
	sqlx.BindDriver("oracle", sqlx.NAMED)
	db, err := sqlx.Open("oracle", dsn)
	if err != nil {
		panic("failed to connect to core policy database")
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(30)

	return db
}

// Get : this use for get one data
func Get(ctx context.Context, db *sqlx.DB, query string, dest interface{}, args map[string]interface{}) error {
	stmt, err := db.PrepareNamedContext(ctx, query)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		return err
	}
	err = stmt.GetContext(ctx, dest, args)
	if err != nil {
		return err
	}
	return nil
}

// Select : this use for get list of data
func Select(ctx context.Context, db *sqlx.DB, query string, dest interface{}, args map[string]interface{}) error {
	stmt, err := db.PrepareNamedContext(ctx, query)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		return err
	}
	err = stmt.SelectContext(ctx, dest, args)
	if err != nil {
		return err
	}
	return nil
}

// Insert : this use for insert/update(id required) data
func Insert(ctx context.Context, db *sqlx.DB, query string, args interface{}) error {

	stmt, err := db.PrepareNamedContext(ctx, query)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

// Executes query, returning result and error
func ExecWithResult(ctx context.Context, db *sqlx.DB, query string, args interface{}) (sql.Result, error) {

	stmt, err := db.PrepareNamedContext(ctx, query)
	defer func() {
		stmt.Close()
	}()
	if err != nil {
		return nil, err
	}
	result, err := stmt.ExecContext(ctx, args)
	if err != nil {
		return nil, err
	}
	return result, err
}
