package database

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    conn := ""
    db, err := sql.Open("postgres", conn)
    return db, err
}
