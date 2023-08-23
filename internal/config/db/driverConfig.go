package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func OpenConnection() (db *sql.DB) {
	godotenv.Load(".env")

	var (
		host     = "localhost"
		port     = 5432
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println(err)
	}

	return db

}
