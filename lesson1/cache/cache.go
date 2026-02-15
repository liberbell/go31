package cache

import "net/http"

type response struct {
	header http.Header
	code   int
	body   []byte
}
