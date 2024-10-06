package service

import (
	"encoding/json"
	"fmt"
	"go-in-memory-cache-api/model"
	"io"
	"log"
	"net/http"
)

type UserService struct{}

func (u *UserService) Get(id int) *model.User {
	var user model.User
	resp, err := http.Get("https://dummyjson.com/user/" + fmt.Sprint(id))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}
