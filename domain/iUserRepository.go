package domain

type IUserRepository interface {
	Create(*User) error
	Get(username string) (*User, error)
}
