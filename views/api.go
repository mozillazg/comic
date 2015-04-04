package views

import (
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/mozillazg/comic/models"
)

func CreateAPIView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c := models.NewComic(
		r.PostFormValue("title"),
		r.PostFormValue("url"),
		r.PostFormValue("description"),
		r.PostFormValue("date"),
		db,
	)
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
	}
	c.Save()
	jsonReponse(w, c, 201)
}

func UpdateAPIView(w http.ResponseWriter, r *http.Request) {
	db, _ := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	n := r.URL.Query().Get(":id")
	id, _ := strconv.ParseInt(n, 10, 64)
	c, err := getComic(w, r, db, id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	c.Title = r.PostFormValue("title")
	c.ImageURL = r.PostFormValue("url")
	c.Description = r.PostFormValue("description")
	c.Date = r.PostFormValue("date")
	c.Save()

	jsonReponse(w, c, 200)
}

func DeleteAPIView(w http.ResponseWriter, r *http.Request) {
	db, _ := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	n := r.URL.Query().Get(":id")
	id, _ := strconv.ParseInt(n, 10, 64)
	c, err := getComic(w, r, db, id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	c.Delete()

	jsonReponse(w, c, 204)
}

func GetAPIView(w http.ResponseWriter, r *http.Request) {
	db, _ := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	n := r.URL.Query().Get(":id")
	id, _ := strconv.ParseInt(n, 10, 64)
	c, err := getComic(w, r, db, id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	jsonReponse(w, c, 200)
}
