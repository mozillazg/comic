package views

import (
	"fmt"
	"net/http"

	"github.com/mozillazg/comic/models"
)

func AtomView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	defer db.Close()

	c, err := models.AllComic(db, "")
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	type data struct {
		Comics     []*models.Comic
		LastUpdate string
	}
	renderTemplate(w, data{c, c[0].Date}, "atom.html", "template/atom.xml")
}
