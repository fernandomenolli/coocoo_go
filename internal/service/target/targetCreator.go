package target

import (
	"coocoo/internal/config/db"
	_ "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Insert() {
	db := db.OpenConnection()

	ins, err := db.Prepare("INSERT INTO target (id, description, url) VALUES (1, 'Linkedin', 'https://linkedin.com/in/fernandomenolli');")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	ins.Exec()

}
