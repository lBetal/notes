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
	Create(userId int, device notes.Device) (int, error)
	GetAll(userId int) ([]notes.Device, error)
	GetById(userId, deviceId int) (notes.Device, error)
	Delete(userId, deviceId int) error
	Update(userId, deviceId int, input notes.UpdateDeviceInput) error
}

type DevicePhoto interface {
	Create(deviceId int, photo notes.Photo) (int, error)
	GetAll(userId, deviceId int) ([]notes.Photo, error)
	GetById(userId, photoId int) (notes.Photo, error)
	Delete(userId, photoId int) error
	Update(userId, photoId int, input notes.UpdateDevicePhotoInput) error
}

type DeviceVideo interface {
	Create(deviceId int, photo notes.Video) (int, error)
	GetAll(userId, deviceId int) ([]notes.Video, error)
	GetById(userId, videoId int) (notes.Video, error)
	Delete(userId, videoId int) error
	Update(userId, videoId int, input notes.UpdateDeviceVideoInput) error
}

type Repository struct {
	Authorization
	Device
	DevicePhoto
	DeviceVideo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Device:        NewDevicePostgres(db),
		DevicePhoto:   NewDevicePhotoPostgres(db),
		DeviceVideo:   NewDeviceVideoPostgres(db),
	}
}
