package main

import (
	"net/http"

	"github.com/barancanatbas/khanmux"
)

func main() {
	r := khanmux.NewRouter()

	r.GET("/name", Home)

	http.ListenAndServe(":8080", r)
}

type User struct {
	Name string `json:"name" xml:"deneme"`
	Age  int    `json:"age" xml:"age"`
}

func Home(c khanmux.Context) error {

	user := User{
		Name: "baran2",
		Age:  20,
	}

	return c.XML(201, user)
}
