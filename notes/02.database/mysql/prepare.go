package mysql

import (
	"context"
	"database/sql"
	"time"
)

func Add(db *sql.DB, people ...Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "INSERT INTO users(name, age) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, person := range people {
		_, err := stmt.Exec(person.Name, person.Age)
		if err != nil {
			return err
		}
	}
	return nil
}
