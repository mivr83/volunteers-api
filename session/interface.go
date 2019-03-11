package session

// holds user session
type User struct {
	DbId uint32
	Role string
}

// session interface defining simple operations which are required by REST api
// this is very simplified version of real session management
type Session interface {
	GetUser(token string) *User
	RemoveUser(token string)
	AddUser(token string, user *User)
}
