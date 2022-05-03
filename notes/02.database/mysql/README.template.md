{
"CreatedAt": "2022-04-28",
"UpdatedAt": "2022-04-29",
"Title": "MySQL - Go notes"
}
---
# MySQL
- Use Query() for SELECT statements, and always handle the returned rows and iterate over them.
- Use Exec() for INSERT, DELETE, UPDATE and other statements that does not return rows.
- Query(query) does not prepare statements, but Query(queryTemplate, params) do.
- Exec() does not prepare statements, but Exec(queryTemplate, params) do.
- Use Prepare() when you want to prepare once and execute many times (many INSERT for example).

## Connection
Make sure to import mysql driver with _.
<<connection.go>>

## Select many rows
<<selectManyRows.go>>

## Select single row
<<selectSingleRow.go>>

## Insert
When inserting we can optionally get affected rows and last inserted id. Notice that
ExecContext(ctx, query, params) does statement prepare under the hood.
<<insert.go>>

## Update
<<update.go>>

## Delete
<<delete.go>>

## Prepare
Prepare statements are useful when you want to prepare the data once and execute many times, for example inserting many rows in table.
<<prepare.go>>

## Transactions
Transactions are very useful when you need to do many database operations (for example creating entries in different tables) and have the possibility to roll back in case one of the operations failed.
<<transactions.go>>

## Null Values
Best is to avoid nullable columns in our table if possible. But if we do have nullable columns
we should use special null types from database/sql or define our own.
