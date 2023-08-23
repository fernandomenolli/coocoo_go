package util

import (
	"coocoo/internal/config/db"
	"coocoo/internal/service/target"
	"log"
)

func FeedDB() {
	feedTable()
	feedData()
}

func feedTable() {
	db := db.OpenConnection()

	_, err := db.Query("CREATE TABLE IF NOT EXISTS target (id int primary key, description varchar(255), url varchar(255));")
	if err != nil {
		panic(err.Error())
	}
}

func feedData() {
	db := db.OpenConnection()

	rows, err := db.Query("SELECT COUNT(*) FROM target;")
	if err != nil {
		panic(err.Error())
	}

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	if count == 0 {
		target.Insert()
	}
}
