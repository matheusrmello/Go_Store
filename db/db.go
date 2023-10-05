package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
func ConnectionDB() *sql.DB {
	conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "postgres", "admin", "go_store_db")
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db
}