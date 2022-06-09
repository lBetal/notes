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

type DevicePhoto interface {
	Create(userId, listId int, item notes.Photo) (int, error)
	GetAll(userId, listId int) ([]notes.Photo, error)
	GetById(userId, itemId int) (notes.Photo, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input notes.UpdateDevicePhotoInput) error
}

type Service struct {
	Authorization
	Device
	DevicePhoto
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Device:        NewDeviceService(repos.Device),
		DevicePhoto:   NewDevicePhotoService(repos.DevicePhoto, repos.Device),
	}
}
