package mysql

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func Connect() (*sql.DB, error) {
	// Create DB connection. sql.DB is not an open connection!
	db, err := sql.Open("mysql", "username:password@tcp(0.0.0.0:3306)/database?parseTime=true")
	if err != nil {
		return nil, err
	}

	defer db.Close()

	// Some connection settings, make sure to set them correct
	db.SetMaxOpenConns(10) // default (if not set) is 0 = unlimited
	db.SetMaxIdleConns(5)  // default (if not set) is 2
	db.SetConnMaxLifetime(time.Minute * 5)

	// In order to make sure we have an available db and correct
	// credentials we need to ping the db, here we ping the db
	// with a 5 seconds timeout limit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Good to go, now is a good time to pass db object to other functions
	// or store it in a settings struct
	return db, nil
}
