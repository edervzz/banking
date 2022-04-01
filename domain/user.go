package domain

type User struct {
	Username       string
	HashedPassword string
	Email          string
}

type UserRepository interface {
	Create(*User) error
}
