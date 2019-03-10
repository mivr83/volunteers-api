package session

type InMemorySession struct {
	sessionStorage map[string]*User
}

func (ims *InMemorySession) GetUser(token string) *User {
	return ims.sessionStorage[token]
}

func (ims *InMemorySession) AddUser(token string, user *User) {
	ims.sessionStorage[token] = user
}

func (ims *InMemorySession) RemoveUser(token string) {
	delete(ims.sessionStorage, token)
}

func Create() *InMemorySession {
	return &InMemorySession{make(map[string]*User)}
}
