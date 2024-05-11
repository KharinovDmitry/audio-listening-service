package model

type Claim struct {
	Title string
	Value any
}

const (
	IDClaimTitle      = "id"
	LoginClaimTitle   = "login"
	RoleClaimTitle    = "role"
	ExpiredClaimTitle = "exp"
)
