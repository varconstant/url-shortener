package main

import (
	"context"
	"net/http"
	"urlShortner/router"
	"urlShortner/server"
	"urlShortner/shortener"
	"urlShortner/store"
	"urlShortner/util"
)

func main() {
	var db store.StorageService = store.NewRedis(context.Background())
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	r := router.NewRouter("url-shortener", mux)
	srv := server.NewServer(shortener.NewUrlShortener(), db)

	r.HandleFunc("/add-url", srv.AddUrl)
	r.HandleFunc("/get-url", srv.FetchUrl)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err = w.Write([]byte("test"))
		util.PrintError(err)
	})
	err = http.ListenAndServe(":9001", mux)
	panic(err)
}
