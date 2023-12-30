package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func ReturnPerson(db *sql.DB, id int) (Person, error) {
	// We set 5 second timeout on our query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p Person
	query := "SELECT id, name, age FROM users WHERE id = ?"
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(&p.Id, &p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return Person{}, fmt.Errorf("person with id %d not found", id)
		}
		return Person{}, err
	}

	return p, nil
}
