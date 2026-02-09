package handlers

import (
	"net/http"
	"Lesson1/user"
	"testing"
	"reflect"
)

func TestBodyToUser(t *testing.T) {
	valid := &user.User{
		ID: bson.NewObjectId(),
		Name: "John",
		Role: "",
	}
	js, err : = json.Marshal(valid)
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
		{
			txt: "empty request body",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
			},
			err: true,
		},
		{
			txt: "malformed data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id": 12}`)),
			},
			u: &user.User{},
			err: true,
		},
		{
			txt: "empty user",
			r: &http.Request{},
			err: true,
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
		if err != nil {
			t.Errorf("Unexpected  error: %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.u, tc.exp) {
			t.Error("Unmarshalled data is different:")
			t.Error(tc.u)
			t.Error(tc.exp)
		}
	}
}
