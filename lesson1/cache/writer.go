package cache

import "net/http"

type Writer struct {
	writer   http.ResponseWriter
	response response
	resource string
}

var (
	_ http.ResponseWriter = (*Writer)(nil)
)

func NewWriter(w http.ResponseWriter, r *http.Request) *Writer {
	return &Writer{
		writer:   w,
		resource: MakeResource(r),
		response: response{
			header: http.Header{},
		},
	}
}

func (w *Writer) Header() http.Header {
	return w.response.header
}

func (w *Writer) WriteHeader(code int) {
	copyHeader(w.response.header, w.writer.Header())
	w.response.code = code
	w.Writer.WriteHeader(code)
}
