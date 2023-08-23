package target

import (
	"coocoo/internal/config/db"
	"coocoo/internal/model"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func selectTargets() []model.Target {
	db := db.OpenConnection()

	selDB, err := db.Query("SELECT * FROM target ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	t := model.Target{}

	res := []model.Target{}

	for selDB.Next() {
		var id int
		var description, url string

		err = selDB.Scan(&id, &description, &url)
		if err != nil {
			panic(err.Error())
		}

		t.Description = description
		t.Url = url

		res = append(res, t)
	}

	return res

}

func Parse() {
	fmt.Println("Started scanning: " + time.Now().String())
	tlist := selectTargets()

	for _, target := range tlist {
		description := target.Description
		url := target.Url
		resp, err := http.Get(url)

		if err != nil {
			panic(err)
		}

		html, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		contentFile, err := ioutil.ReadFile(description + ".txt")

		if string(html) != string(contentFile) {
			fileWriter(description, html)
		}

	}
	fmt.Println("Finished scanning")
}

func fileWriter(description string, html []byte) {
	data, err := os.Create(description + ".txt")
	if err != nil {
		panic(err)
	}

	io.Copy(data, strings.NewReader(string(html)))
	fmt.Println("-- Content Updated")
}
