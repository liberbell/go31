package user

type User struct {
	ID   bson.ObjectId `json: "id"`
	Nmae string        `json: "name"`
	Role string        `json: "role"`
}
