package main

import (
	"bytes"
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/bmizerany/pat"
	"github.com/mozillazg/comic/views"
)

func BasicAuth(
	f func(http.ResponseWriter, *http.Request), user, pass []byte,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const basicAuthPrefix string = "Basic "

		// Get the Basic Authentication credentials
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			// Check credentials
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(basicAuthPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && bytes.Equal(pair[0], user) &&
					bytes.Equal(pair[1], pass) {
					// Delegate request to the given handle
					f(w, r)
					return
				}
			}
		}

		// Request Basic Authentication otherwise
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized),
			http.StatusUnauthorized)
	}
}

func main() {
	listen := ":8080"
	user := []byte("user")
	pass := []byte("passwd")

	router := pat.New()
	router.Get("/", http.HandlerFunc(views.LastComicView))
	router.Get("/first", http.HandlerFunc(views.FirstComicView))
	router.Get("/last", http.HandlerFunc(views.LastComicView))
	router.Get("/random", http.HandlerFunc(views.RandomComicView))
	router.Get("/admin", http.HandlerFunc(BasicAuth(views.ListView, user, pass)))

	router.Post("/api/comics", http.HandlerFunc(BasicAuth(views.CreateAPIView, user, pass)))
	router.Get("/api/comics/:id", http.HandlerFunc(BasicAuth(views.GetAPIView, user, pass)))
	router.Del("/api/comics/:id", http.HandlerFunc(BasicAuth(views.DeleteAPIView, user, pass)))
	router.Put("/api/comics/:id", http.HandlerFunc(BasicAuth(views.UpdateAPIView, user, pass)))

	router.Get("/:id", http.HandlerFunc(views.GetComicView))

	http.Handle("/", router)
	log.Printf("listen %s\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
