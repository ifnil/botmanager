package database

import (
	"database/sql"

	"github.com/spf13/viper"
	_ "modernc.org/sqlite"
)

type Database struct {
	database *sql.DB
	taskchan chan *sql.Stmt
}

func New() *Database {
	db, err := sql.Open("sqlite", viper.GetString("database.path"))
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		return nil
	}

	return &Database{
		database: db,
		taskchan: make(chan *sql.Stmt),
	}
}

func (db *Database) Init() error {
	if err := db.database.Ping(); err != nil {
		close(db.taskchan)
		return err
	}

	return db.createTables()
}

func (db *Database) DB() *sql.DB {
	return db.database
}
