package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID   bson.ObjectId `json: "id" storm: "id"`
	Nmae string        `json: "name"`
	Role string        `json: "role"`
}

const {
	dbPath = "users.db"
}

func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}