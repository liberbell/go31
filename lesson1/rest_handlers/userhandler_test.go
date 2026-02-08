package handlers

import (
	"net/http"
	"os/user"
	"testing"
)

func TestBodyToUser(t *testing.T) {
	ts := []struct {
		txt string
		r   *http.Request
		u   *user.User
	}
}
