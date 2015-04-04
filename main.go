package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/mozillazg/comic/views"
)

func main() {
	listen := ":8080"

	router := pat.New()
	router.Get("/", http.HandlerFunc(views.LastComicView))
	router.Get("/first", http.HandlerFunc(views.FirstComicView))
	router.Get("/last", http.HandlerFunc(views.LastComicView))
	router.Get("/random", http.HandlerFunc(views.RandomComicView))
	router.Get("/admin", http.HandlerFunc(views.ListView))

	router.Post("/api/comics", http.HandlerFunc(views.CreateAPIView))
	router.Get("/api/comics/:id", http.HandlerFunc(views.GetAPIView))
	router.Del("/api/comics/:id", http.HandlerFunc(views.DeleteAPIView))
	router.Put("/api/comics/:id", http.HandlerFunc(views.UpdateAPIView))

	router.Get("/:id", http.HandlerFunc(views.GetComicView))

	http.Handle("/", router)
	log.Printf("listen %s\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
