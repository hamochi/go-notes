package main

import (
	"context"
	"database/sql"
	"time"
)

func NewUser(db *sql.DB, p Person, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err == nil {
		return err
	}

	// If we return early due to error, we make sure to roll back everything
	defer tx.Rollback()

	firstQuery := "INSERT INTO users(name, age) VALUES(?,?)"
	res, err := tx.ExecContext(ctx, firstQuery, p.Name, p.Age)
	if err == nil {
		return err
	}

	// Get lastInsertId so that we can use ut for the next insert
	userId, err := res.LastInsertId()
	if err != nil {
		return err
	}

	secondQuery := "INSERT INTO account(user_id, email) VALUES(?,?)"
	_, err = tx.ExecContext(ctx, secondQuery, userId, email)
	if err == nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
