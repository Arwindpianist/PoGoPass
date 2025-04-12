package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenEncryptedDB(path string, key []byte) (*sql.DB, error) {
	// Convert the key to a string representation if necessary
	connStr := fmt.Sprintf("%s?_pragma_key=%s&_pragma_cipher_page_size=4096", path, string(key))
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, err
	}

	// Create table if not exists
	query := `
	CREATE TABLE IF NOT EXISTS passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		service TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	)`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
