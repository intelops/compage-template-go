package service

import (
	"errors"
	"github.com/kube-tarian/compage-template-go/pkg/models"
)

var users = make(map[string]models.User)

type UserService struct {
}

func (userService UserService) CreateUser(user models.User) error {
	users[user.Id] = user

	return nil
}

func (userService UserService) UpdateUser(id string, user models.User) error {
	if id != user.Id {
		return errors.New("id and payload don't match")
	}
	users[user.Id] = user

	return nil
}

func (userService UserService) DeleteUser(id string) error {
	if _, ok := users[id]; ok {
		delete(users, id)
		return nil
	}

	return errors.New("user not found")
}

func (userService UserService) ListUsers() ([]models.User, error) {
	v := make([]models.User, len(users))
	for _, value := range users {
		v = append(v, value)
	}

	return v, nil
}

func (userService UserService) GetUser(id string) (models.User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}

	return models.User{}, errors.New("user not found")
}
