package user

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestCRUD(t *testing.T) {
	t.Log("Create")
	u := &User{
		ID: bson.NewObjectId(),
	}
}
