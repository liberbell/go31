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
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := BodyToUser(tc.r, tc.u)
		if tc.err {
			if err == nil {
				t.Error("Expected error, got none")
			}
			continue
		}
	}
}
