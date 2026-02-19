package handlers

type mockWriter response

func newMockWriter() *mockWriter {
	return &mockWriter{
		body:   []body{},
		header: http.header{},
	}
}
