package main

import (
	"bytes"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bmizerany/pat"
	"github.com/mozillazg/comic/views"
)

type ViewFunc func(http.ResponseWriter, *http.Request)

func BasicAuth(f ViewFunc, user, passwd []byte) ViewFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basicAuthPrefix := "Basic "

		// 获取 request header
		auth := r.Header.Get("Authorization")
		// 如果是 http basic auth
		if strings.HasPrefix(auth, basicAuthPrefix) {
			// 解码认证信息
			payload, err := base64.StdEncoding.DecodeString(
				auth[len(basicAuthPrefix):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && bytes.Equal(pair[0], user) &&
					bytes.Equal(pair[1], passwd) {
					// 执行被装饰的函数
					f(w, r)
					return
				}
			}
		}

		// 认证失败，提示 401 Unauthorized
		// Restricted 可以改成其他的值，作用类似于 session ,这样就不会每次访问页面都提示登录
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		// 401 状态码
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func main() {
	listen := os.Getenv("COMIC_LISTEN")
	user := []byte(os.Getenv("COMIC_USER"))
	pass := []byte(os.Getenv("COMIC_PASSWD"))

	router := pat.New()
	router.Get("/", http.HandlerFunc(views.LastComicView))
	router.Get("/first", http.HandlerFunc(views.FirstComicView))
	router.Get("/last", http.HandlerFunc(views.LastComicView))
	router.Get("/random", http.HandlerFunc(views.RandomComicView))
	router.Get("/admin", http.HandlerFunc(BasicAuth(views.ListView, user, pass)))

	router.Post("/api/comics", http.HandlerFunc(BasicAuth(views.CreateAPIView, user, pass)))
	router.Get("/api/comics", http.HandlerFunc(BasicAuth(views.ListAPIView, user, pass)))
	router.Get("/api/comics/:id", http.HandlerFunc(BasicAuth(views.GetAPIView, user, pass)))
	router.Del("/api/comics/:id", http.HandlerFunc(BasicAuth(views.DeleteAPIView, user, pass)))
	router.Put("/api/comics/:id", http.HandlerFunc(BasicAuth(views.UpdateAPIView, user, pass)))

	router.Get("/archive", http.HandlerFunc(views.ArchiveView))
	router.Get("/atom", http.HandlerFunc(views.AtomView))
	router.Get("/:id", http.HandlerFunc(views.GetComicView))

	http.Handle("/", router)
	log.Printf("listen %s\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
