package cache

import (
	"net/http"
	"net/url"
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

	res := "/test/url?with=params"
	u, err := url.Parse(res)
	if err != nil {
		t.Fatal("Invalid url")
	}

	req := &http.Request{
		URL: u,
	}
	t.Log("test NewWriter")
	w := NewWriter(mw, req)
	if w.resource != res {
		t.Errorf("Resources are different, Expected: %s / Actual: %s", res, w.resource)
	}

	if w.writer != mw {
		t.Fatal("Writer not assigned")
	}

	t.Log("test Header")
	h := w.Header()
}
