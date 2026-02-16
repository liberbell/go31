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
