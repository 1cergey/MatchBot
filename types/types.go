package types

type Player struct {
	UserID	  string `db:"user_id"`
	UserName  string `db:"username"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type Play struct {
	ChatID  int64
	Players []Player
}
