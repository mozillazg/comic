package views

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mozillazg/comic/models"
)

func ListView(w http.ResponseWriter, r *http.Request) {
	db, err := models.NewConnect(dbPath)
	db.Begin()
	defer db.Close()

	c, err := models.AllComic(db)
	if err != nil {
		fmt.Printf("%p", err)
		http.Error(w, http.StatusText(404), 404)
	}
	type data struct {
		Comics []*models.Comic
	}
	fmt.Println(c)
	renderTemplate(w, data{c}, "admin.tpl", "template/admin.tpl")
}
