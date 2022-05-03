package mysql

import (
	"context"
	"database/sql"
	"time"
)

func AddPerson(db *sql.DB, name string, age int) (int64, int64, error) {
	// We set 5 second timeout on our query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO users(name, age) VALUES(?, ?)"
	res, err := db.ExecContext(ctx, query, name, age)
	if err != nil {
		return 0, 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	return rowsAffected, lastId, nil
}
