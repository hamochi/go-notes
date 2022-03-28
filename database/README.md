# MySQL
- Use Query() for SELECT statements, and always handle the returned rows and iterate over them.
- Use Exec() for INSERT, DELETE, UPDATE and other statements that does not return rows.
- Query(query) does not prepare statements, but Query(queryTemplate, params) do.
- Exec() does not prepare statements, but Exec(queryTemplate, params) do.
- Use Prepare() when you want to prepare once and execute many times (many INSERT for example).

## Connection
Make sure to import mysql driver with _.
```go
package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create DB connection. sql.DB is not an open connection!
	db, err := sql.Open("mysql", "username:password@tcp(0.0.0.0:3306)/database?parseTime=true")
	if err != nil {
		panic(err)
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
		panic(err)
	}

	// Good to go, now is a good time to pass db object to other functions
	// or store it in a settings struct
}
```

## Select many rows
```go
type Person struct {
	Id   int
	Name string
	Age  int
}

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
```

## Select single row
```go
type Person struct {
	Id   int
	Name string
	Age  int
}

func ReturnPerson(db *sql.DB, id int) (Person, error) {
	// We set 5 second timeout on our query
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var p Person
	query := "SELECT id, name, age FROM users WHERE id = ?"
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(&p.Id, &p.Name, &p.Age)
	if err != nil {
		return Person{}, err
	}

	return p, nil
}
```
## Insert
When inserting we can optionally get affected rows and last inserted id. Notice that
ExecContext(ctx, query, params) does statement prepare under the hood.
```go
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
```
## Update
```go
type Person struct {
	Id   int
	Name string
	Age  int
}

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
```
## Delete
```go
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
```

## Prepare
```go
type Person struct {
    Id   int
    Name string
    Age  int
}

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
```
## Transactions
```go
type Person struct {
	Id   int
	Name string
	Age  int
}

func NewUser(db *sql.DB, p Person, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err == nil {
		return err
	}

	// if we return early due to error, we make sure to roll back everything
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
```

## Null Values
Best is to avoid nullable columns in our table if possible. But if we do have nullable columns
we should use special null types from database/sql or define our own.