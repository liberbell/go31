package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID   bson.ObjectId `json: "id" storm: "id"`
	Nmae string        `json: "name"`
	Role string        `json: "role"`
}
