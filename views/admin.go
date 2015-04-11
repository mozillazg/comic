package views

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mozillazg/comic/models"
)

func ListView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
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
	renderTemplate(w, data{c}, "admin.html", "template/admin.html")
}
