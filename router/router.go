package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	name string
	mux  *http.ServeMux
}

func (router *Router) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request)) {
	fullPath := fmt.Sprintf("/%s%s", router.name, path)
	router.mux.HandleFunc(fullPath, f)
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	router.mux.ServeHTTP(writer, request)
}

func NewRouter(name string, mux *http.ServeMux) *Router {
	return &Router{
		name: name,
		mux:  mux,
	}
}
