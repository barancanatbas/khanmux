# khanmux

#### An http router I made to improve myself. May have many errors and shortcomings
#### Basically, it separates the discarded requests into its methods.

## a small sample

```go
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
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func Home(c khanmux.Context) error {

	user := User{
		Name: "baran2",
		Age:  20,
	}

	return c.XML(201, user)
}

```
