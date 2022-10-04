package infrastructure

import (
	"errors"
	"login-go/domain"
	"sort"
)

type UserRepositoryStub struct {
	users []domain.User
}

func (r *UserRepositoryStub) Create(user *domain.User) error {
	r.users = append(r.users, *user)
	return nil
}

func (r *UserRepositoryStub) Get(username string) (*domain.User, error) {
	if len(r.users) == 0 {
		return nil, errors.New("username not found")
	}

	idx := sort.Search(len(r.users), func(i int) bool {
		v := r.users[i].Username == username
		return v
	})
	if idx == len(r.users) {
		return nil, errors.New("username not found")
	}
	return &r.users[idx], nil
}

func NewUserRepositoryStub() *UserRepositoryStub {
	return &UserRepositoryStub{}
}
