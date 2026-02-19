package handlers

import "net/http"

type response struct {
	header http.Header
	code   int
	body   []byte
}

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
