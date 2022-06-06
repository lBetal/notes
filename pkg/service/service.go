package service

import (
	"github.com/lBetal/todo"
	"github.com/lBetal/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Device interface {
	Create(userId int, list todo.Device) (int, error)
	GetAll(userId int) ([]todo.Device, error)
	GetById(userId, listId int) (todo.Device, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateDeviceInput) error
}

type DeviceItem interface {
	Create(userId, listId int, item todo.DeviceItem) (int, error)
	GetAll(userId, listId int) ([]todo.DeviceItem, error)
	GetById(userId, itemId int) (todo.DeviceItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateDeviceItemInput) error
}

type Service struct {
	Authorization
	Device
	DeviceItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Device:        NewDeviceService(repos.Device),
		DeviceItem:    NewDeviceItemService(repos.DeviceItem, repos.Device),
	}
}
