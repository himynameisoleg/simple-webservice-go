package models

type User struct {
	ID        int
	LastName  string
	FirstName string
}

var (
	users  []*User
	nextID = 1
)
