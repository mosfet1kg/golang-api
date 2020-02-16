package main

import (
	"fmt"
	"github.com/mosfet1kg/golang-api/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
