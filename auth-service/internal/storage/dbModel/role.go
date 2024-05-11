package dbModel

type Role struct {
	ID   int    `db:"id"`
	Name string `db:"role"`
}
