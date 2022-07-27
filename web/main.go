package main

import (
	"context"
	"fmt"
	"net/http"
)

type User struct {
	ID int
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World") // Hello, Worldってアクセスした人に返すよ！
	var body struct {
		ID int
	}
	u, err := GetUser(r.Context, body.ID)
}

func GetUser(ctx context.Context, id int) (*User, error) {

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
