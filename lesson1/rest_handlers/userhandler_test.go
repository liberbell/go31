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
		err bool
		exp *user.User
	}{
		{
			txt : "nil request",
			err : true,
		}
	}
}
