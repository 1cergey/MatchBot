package types

type Player struct {
	UserName  string `db:"username"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"` 
}
