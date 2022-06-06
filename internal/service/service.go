package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Device interface {
	Create(userId int, list notes.Device) (int, error)
	GetAll(userId int) ([]notes.Device, error)
	GetById(userId, listId int) (notes.Device, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input notes.UpdateDeviceInput) error
}

type DeviceItem interface {
	Create(userId, listId int, item notes.DeviceItem) (int, error)
	GetAll(userId, listId int) ([]notes.DeviceItem, error)
	GetById(userId, itemId int) (notes.DeviceItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input notes.UpdateDeviceItemInput) error
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
