package main

import (
	"coocoo/internal/service/target"
	"coocoo/internal/util"
	"time"
)

func main() {
	util.FeedDB()

	scanJob()
	forever()
}

func scanJob() {
	target.Parse()
}

func forever() {
	for range time.Tick(time.Second * 5) {
		scanJob()
	}
}
