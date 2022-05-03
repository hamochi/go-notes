package mysql

import (
	"context"
	"database/sql"
	"time"
)

func ReturnPeople(db *sql.DB, name string) ([]Person, error) {
	// We set 5 second timeout on our query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use db.Query() if you don't want to pass context
	query := "SELECT id, name, age FROM users WHERE name = ?"
	rows, err := db.QueryContext(ctx, query, name)
	if err != nil {
		return nil, err
	}

	// rows.Close will be called when rows.Next() reaches it end,
	// but if we return earlier due to error, better make sure we close
	defer rows.Close()

	var people []Person
	for rows.Next() {
		p := Person{}
		err := rows.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return people, nil
}
