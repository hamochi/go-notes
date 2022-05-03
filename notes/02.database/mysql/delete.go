package mysql

import (
	"context"
	"database/sql"
	"time"
)

func DeletePerson(db *sql.DB, id int) (int64, error) {
	// We set 5 second timeout on our query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE from users WHERE id = ?"
	res, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
