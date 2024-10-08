package handler

import (
	"encoding/json"
	"go-in-memory-cache-api/pkg/cache"
	"go-in-memory-cache-api/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var userService service.UserService

type UserHandler struct{}

var myCache = cache.GetCache(time.Second)

func (u *UserHandler) Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		user := userService.Get(id)
		myCache.Set(cache.UserCacheKey(user.ID), *user)
		bytes, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(bytes)
	}
}
