package main

import (
	"context"
	"database/sql"
	"time"
)

func UpdatePerson(db *sql.DB, p Person) (int64, error) {
	// We set 5 second timeout on our query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "UPDATE users SET name = ?, age = ?, WHERE id = ?"
	res, err := db.ExecContext(ctx, query, p.Name, p.Age, p.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
