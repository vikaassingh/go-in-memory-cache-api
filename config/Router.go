package config

import (
	"encoding/json"
	"fmt"
	"go-in-memory-cache-api/handler"
	"go-in-memory-cache-api/pkg/cache"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var userHandler handler.UserHandler
var myCache = cache.GetCache(time.Minute)

func IntializeRouter() {

	r := mux.NewRouter()

	r.Use(RequestMiddleware)

	r.HandleFunc("/api/hello-world", HelloHandler).Methods(http.MethodGet)

	// userRoutes := r.PathPrefix("/api/user").Subrouter()
	// userRoutes.Use(CheckUserCacheMiddleware)
	// userRoutes.HandleFunc("/{id}", userHandler.Get()).Methods(http.MethodGet)
	r.Handle("/api/user/{id}", CheckUserCacheMiddleware(userHandler.Get())).Methods(http.MethodGet)

	fmt.Println("server listen on 127.0.0.1:8000")
	http.ListenAndServe(":8000", r)
}

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// log.Printf("Request: %s %s", r.Method, r, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func CheckUserCacheMiddleware(next func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("cache:%p-%v", myCache, myCache)
		idStr := mux.Vars(r)["id"]
		var userCacheKey cache.UserCacheKey
		err := json.Unmarshal([]byte(idStr), &userCacheKey)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		userCacheNode, ok := myCache.Get(userCacheKey)
		if ok {
			// fmt.Printf("\nuserCacheNode:%v\n", userCacheNode)
			bytes, err := json.Marshal(userCacheNode.User)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(bytes)
			return
		}
		next(w, r)
	})
}
