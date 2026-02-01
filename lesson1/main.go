package main

import (
	"fmt"
	"lesson1/rest_handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rest_handlers.RootHandler)
	http.HandleFunc("/users", rest_handlers.UserHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
