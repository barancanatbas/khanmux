package main

import (
	"github.com/barancanatbas/khanmux"
)

func main() {
	r := khanmux.NewRouter()

	r.GET("/name", Home)
	r.POST("/user", SaveUser)

	r.Run(":8080")
}

type User struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func SaveUser(c khanmux.Context) error {
	var user User
	err := c.Find(&user)
	if err != nil {
		return c.JSON(400, "Veri yok")
	}

	return c.XML(200, user)
}

func Home(c khanmux.Context) error {

	user := User{
		Name: "baran2",
		Age:  20,
	}

	return c.JSON(201, user)
}
