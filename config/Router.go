package config

import (
	"fmt"
	"go-in-memory-cache-api/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var userHandler handler.UserHandler

func IntializeRouter() {

	r := mux.NewRouter()

	r.Use(RequestMiddleware)

	r.HandleFunc("/api/hello-world", HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/user/{id}", userHandler.Get()).Methods(http.MethodGet)
	fmt.Println("server listen on 127.0.0.1:8000")
	http.ListenAndServe(":8000", r)
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
