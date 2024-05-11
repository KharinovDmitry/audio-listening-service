package model

var (
	ListenerRole = "Listener"
	ArtistRole   = "Artist"
)

type Role struct {
	ID   int
	Name string
}
