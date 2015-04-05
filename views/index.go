package views

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mozillazg/comic/models"
)

var tplPath = "template/index.html"
var tplName = "index.html"

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
		http.NotFound(w, r)
		return
	}
	c, err := models.GetComic(db, id)
	if err != nil {
		fmt.Printf("%p", err)
		http.NotFound(w, r)
		return
	}
	prev, _ := models.PrevComic(db, id)
	next, _ := models.NextComic(db, id)
	if prev == nil {
		prev = &models.Comic{}
	}
	if next == nil {
		next = &models.Comic{}
	}
	data := struct {
		Comic, Prev, Next *models.Comic
	}{c, prev, next}
	renderTemplate(w, data, tplName, tplPath)
}

func FirstComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.FirstComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.NotFound(w, r)
		return
	}
	url := urlFor(r, "/"+strconv.FormatInt(c.ID, 10))
	http.Redirect(w, r, url, 302)
}

func LastComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.LastComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.NotFound(w, r)
		return
	}
	url := urlFor(r, "/"+strconv.FormatInt(c.ID, 10))
	http.Redirect(w, r, url, 302)
}

func RandomComicView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.RandomComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.NotFound(w, r)
		return
	}
	url := urlFor(r, "/"+strconv.FormatInt(c.ID, 10))
	http.Redirect(w, r, url, 302)
}

func ArchiveView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.AllComic(db, "")
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	type data struct {
		Comics []*models.Comic
	}
	renderTemplate(w, data{c}, "archive.html", "template/archive.html")
}
