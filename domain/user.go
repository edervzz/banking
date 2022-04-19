package domain

type User struct {
	Username       string
	HashedPassword string
	Email          string
	Role           string
}

type UserRepository interface {
	Create(*User) error
	Find(*User) error
}
