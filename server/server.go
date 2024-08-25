package server

import (
	"net/http"
	"urlShortner/shortener"
	"urlShortner/store"
	"urlShortner/util"
)

type Server struct {
	Shortener shortener.Shortener
	Db        store.StorageService
}

func NewServer(shortener shortener.Shortener, db store.StorageService) *Server {
	return &Server{
		Shortener: shortener,
		Db:        db,
	}
}

func (s *Server) AddUrl(w http.ResponseWriter, r *http.Request) {
	var url, userId string
	queryMap := r.URL.Query()
	if url = queryMap.Get("url"); url == "" {
		_, err := w.Write([]byte("URL not passed in request param"))
		util.PrintError(err)
		return
	}
	if userId = queryMap.Get("userId"); userId == "" {
		_, err := w.Write([]byte("userId not provided"))
		util.PrintError(err)
		return
	}
	shorten, err := s.Shortener.Shorten(shortener.ShortenRequest{Url: url, UserId: userId})
	if err != nil {
		_, err = w.Write([]byte("error: " + err.Error()))
		util.PrintError(err)
		return
	}
	err = s.Db.Save(shorten, url)
	if err != nil {
		util.PrintError(err)
		return
	}
	_, err = w.Write([]byte(shorten))
	util.PrintError(err)
}

func (s *Server) FetchUrl(w http.ResponseWriter, r *http.Request) {
	var shortened, userId string
	queryMap := r.URL.Query()
	if userId = queryMap.Get("userId"); userId == "" {
		_, err := w.Write([]byte("userId not provided"))
		util.PrintError(err)
		return
	}
	if shortened = queryMap.Get("shortened"); shortened == "" {
		_, err := w.Write([]byte("shortened not passed in request param"))
		util.PrintError(err)
		return
	}

	fetch, err := s.Db.Fetch(shortened)
	if err != nil {
		util.PrintError(err)
		return
	}
	_, err = w.Write([]byte(fetch.(string)))
	if err != nil {
		util.PrintError(err)
		return
	}
}
