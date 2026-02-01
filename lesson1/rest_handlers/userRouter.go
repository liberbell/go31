package handlers

import (
	"fmt"
	"net/http"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}
