package views

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
	"github.com/mozillazg/comic/models"
)

var dbPath = "user=ucomic password=comic dbname=comic"

func getComic(w http.ResponseWriter, r *http.Request, db *sql.DB, n int64) (
	c *models.Comic, err error,
) {
	c, err = models.GetComic(db, n)
	if err != nil {
		fmt.Printf("%p", err)
		http.NotFound(w, r)
		return
	}
	return
}

func renderTemplate(w http.ResponseWriter, c interface{}, name string, path string) {
	b, _ := ioutil.ReadFile(path)
	t, _ := template.New(name).Parse(string(b))
	err := t.Execute(w, c)
	log.Println(err)
}

func jsonReponse(w http.ResponseWriter, v interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if statusCode != 204 {
		b, _ := json.Marshal(v)
		fmt.Fprint(w, string(b))
	}
}

func urlFor(r *http.Request, uri string) (url string) {
	url = "http://" + r.Host + uri
	log.Print(url)
	return
}
