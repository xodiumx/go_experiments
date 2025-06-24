package repo

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
	Delete(id int) error
}
