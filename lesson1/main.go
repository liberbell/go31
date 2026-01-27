package main

import "net/http"

func main() {
	http.ListenAndServe("localhost:11111", nil)
}
