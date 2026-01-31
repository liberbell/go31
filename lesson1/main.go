package main

import (
	"fmt"
	"handlers/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
