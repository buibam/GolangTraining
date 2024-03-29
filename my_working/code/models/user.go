package models

type User struct {
	ID        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

func GetUser() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
