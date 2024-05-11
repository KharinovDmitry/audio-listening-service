package model

type User struct {
	ID       int
	Login    string
	Password string
	Role     Role
}
