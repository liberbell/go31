package cache

import (
	"net/http"
	"testing"
)

type mockWriter response

func newMockWriter() *mockWriter {
	return &mockWriter{
		body:   []byte{},
		header: http.Header{},
	}
}

func (mw *mockWriter) Write(b []byte) (int, error) {
	mw.body = make([]byte, len(b))
	for k, v := range b {
		mw.body[k] = v
	}
	return len(b), nil
}

func (mw *mockWriter) WriteHeader(code int) {
	mw.code = code
}

func (mw *mockWriter) Header() http.Header {
	return mw.header
}

func TestWriter(t *testing.T) {
	mw := newMockWriter()
}
