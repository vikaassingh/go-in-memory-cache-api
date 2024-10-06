package handler

import (
	"encoding/json"
	"go-in-memory-cache-api/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var userService service.UserService

type UserHandler struct{}

func (u *UserHandler) Get() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		bytes, err := json.Marshal(userService.Get(id))
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write(bytes)
	}
}
