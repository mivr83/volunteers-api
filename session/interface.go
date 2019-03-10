package session

type User struct {
	DbId uint32
	Role string
}

type Session interface {
	GetUser(token string) *User
	RemoveUser(token string)
	AddUser(token string, user *User)
}
