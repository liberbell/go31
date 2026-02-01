package main

import (
	"fmt"
	"net/http"
	"os"
)

func usersRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1...\n"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/users", usersRouter)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
