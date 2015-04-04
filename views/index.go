package views

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mozillazg/comic/models"
)

func IndexView(w http.ResponseWriter, r *http.Request) {
	LastComicView(w, r)
}

func GetComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	n := r.URL.Query().Get(":id")
	id, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	c, err := models.GetComic(db, id)
	renderTemplate(w, c, "index.tpl", "template/index.tpl")
}

func FirstComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.FirstComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	renderTemplate(w, c, "index.tpl", "template/index.tpl")
}

func LastComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.LastComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	renderTemplate(w, c, "index.tpl", "template/index.tpl")
}

func RandomComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.RandomComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
	}
	renderTemplate(w, c, "index.tpl", "template/index.tpl")
}
