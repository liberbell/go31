package user

import (
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestCRUD(t *testing.T) {
	t.Log("Create")
	u := &User{
		ID:   bson.NewObjectId(),
		Name: "John",
		Role: "Tester",
	}
	err := u.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	t.Log("Read")
	u2, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record: %s", err)
	}
	if !reflect.DeepEqual(u2, u)
}
