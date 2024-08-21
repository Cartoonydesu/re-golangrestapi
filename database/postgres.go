package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgres() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Panic(err)
	}
	return db
}

func ResetDB() {
	db := NewPostgres()
	defer db.Close()
	_, err := db.Exec("DROP TABLE if exists skill CASCADE;")
	if err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS skill (
			key TEXT PRIMARY KEY,
			name TEXT NOT NULL DEFAULT '',
			description TEXT NOT NULL DEFAULT '',
			logo TEXT NOT NULL DEFAULT '',
			tags TEXT [] NOT NULL DEFAULT '{}'
		);
		INSERT INTO skill (key, name, description, logo, tags)
		VALUES (
			'go',
			'Go',
			'Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.',
			'https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg',
			'{go, golang}'
		),
		(
			'nodejs',
			'Node.js',
			'Node.js is an open-source, cross-platform, JavaScript runtime environment that executes JavaScript code outside of a browser.',
			'https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg',
			'{runtime, javascript}')
			`)
	if err != nil {
		log.Panic(err)
	}
}
