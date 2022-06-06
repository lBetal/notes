package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lBetal/notes"
)

type Authorization interface {
	CreateUser(user notes.User) (int, error)
	GetUser(username, password string) (notes.User, error)
}

type Device interface {
	Create(userId int, list notes.Device) (int, error)
	GetAll(userId int) ([]notes.Device, error)
	GetById(userId, listId int) (notes.Device, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input notes.UpdateDeviceInput) error
}

type DeviceItem interface {
	Create(listId int, item notes.DeviceItem) (int, error)
	GetAll(userId, listId int) ([]notes.DeviceItem, error)
	GetById(userId, itemId int) (notes.DeviceItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input notes.UpdateDeviceItemInput) error
}

type Repository struct {
	Authorization
	Device
	DeviceItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Device:        NewDevicePostgres(db),
		DeviceItem:    NewDeviceItemPostgres(db),
	}
}
