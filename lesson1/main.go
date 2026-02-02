package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func postError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func usersRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			return
		case http.MethodPost:
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}
	path = strings.TrimPrefix(path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
	}
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
	http.HandleFunc("/users/", usersRouter)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
