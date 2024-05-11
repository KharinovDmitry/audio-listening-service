package dbModel

type User struct {
	ID       int    `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
	RoleID   string `db:"role_id"`
}
