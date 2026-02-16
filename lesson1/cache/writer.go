package cache

import "net/http"

type Writer struct {
	writer   http.ResponseWriter
	response response
	resource string
}
