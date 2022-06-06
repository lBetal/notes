package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lBetal/todo"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type Device interface {
	Create(userId int, list todo.Device) (int, error)
	GetAll(userId int) ([]todo.Device, error)
	GetById(userId, listId int) (todo.Device, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateDeviceInput) error
}

type DeviceItem interface {
	Create(listId int, item todo.DeviceItem) (int, error)
	GetAll(userId, listId int) ([]todo.DeviceItem, error)
	GetById(userId, itemId int) (todo.DeviceItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateDeviceItemInput) error
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
