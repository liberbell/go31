package main

import (
	"fmt"
	"lesson1/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/users", handlers.UserHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
