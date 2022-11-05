package service

import (
	"github.com/kube-tarian/compage-template-go/pkg/models"
	"github.com/kube-tarian/compage-template-go/pkg/rest/server/dao"
)

type UserService struct {
}

var userDao = dao.UserDao{}

func (userService UserService) CreateUser(user models.User) error {
	return userDao.CreateUser(user)
}

func (userService UserService) UpdateUser(id string, user models.User) error {
	return userDao.UpdateUser(id, user)
}

func (userService UserService) DeleteUser(id string) error {
	return userDao.DeleteUser(id)
}

func (userService UserService) ListUsers() ([]models.User, error) {
	return userDao.ListUsers()
}

func (userService UserService) GetUser(id string) (models.User, error) {
	return userDao.GetUser(id)
}
