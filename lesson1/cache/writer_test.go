package cache

type mockWriter response

func newMockWriter() *mockWriter {
	return &mockWriter{
		body: []byte{},
	}
}
