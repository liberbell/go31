package cache

import "net/http"

type mockWriter response

func newMockWriter() *mockWriter {
	return &mockWriter{
		body:   []byte{},
		header: http.Header{},
	}
}
