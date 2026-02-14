package user

import (
	"reflect"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Remove(dbPath)
}

func BenchmarkCRUD(b *benchmark.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := &User{
			ID: bson.NewObjectId(),
			Name: "John",
			Role: "Tester",
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving a record: %s", err)
		}
		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Error sretrieving a record: %s", err)
		}

		u.Role = "developer"
		err = u.Save()
		if err != nil {
			b.Fatalf("Error saving a record: %s", err)
		}
	}

	t.Log("Create")
	u := &User{
		ID: bson.NewObjectID(),
		Name: "John",
		Role: "Tester"
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
	if !reflect.DeepEqual(u2, u) {
		t.Error("Records do not match")
	}
	u.Role = "developer"
	err = u.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	u3, err := One(u.ID)
	t.Log("Delete")
	err = Delete(u.ID)
	if err != nil {
		t.
	}
}

#test
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
	if !reflect.DeepEqual(u2, u) {
		t.Error("Records do not match")
	}
	t.Log("Update")
	u.Role = "developer"
	err = u.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	u3, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record: %s", err)
	}
	if !reflect.DeepEqual(u3, u) {
		t.Error("Records do not match")
	}
	t.Log("Delete")
	err = Delete(u.ID)
	if err != nil {
		t.Fatalf("Error removing a record: %s", err)
	}
	_, err = One(u.ID)
	if err != nil {
		t.Fatal("Record shoud not exist anymore")
	}
	if err != storm.ErrNotFound {
		t.Fatalf("Error retrieving non-existing record: %s", err)
	}

	t.Log("Read all")
	u2.ID = bson.NewObjectId()
	u3.ID = bson.NewObjectID()
	err = u.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	err = u2.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	err = u3.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}

	users, err := All()
	if err != nil {
		t.Fatalf("Error reading all records: %s", err)
	}
	if len(users) != 3 {
		t.Errorf("Different number of records retrieved. Expected 3 / Actual: %d", len(users))
	}
}
