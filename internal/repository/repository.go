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

type DeviceAudio interface {
	Create(deviceId int, photo notes.Audio) (int, error)
	GetAll(userId, deviceId int) ([]notes.Audio, error)
	GetById(userId, audioId int) (notes.Audio, error)
	Delete(userId, audioId int) error
	Update(userId, audioId int, input notes.UpdateDeviceAudioInput) error
}

type DeviceMessage interface {
	Create(deviceId int, photo notes.Message) (int, error)
	GetAll(userId, deviceId int) ([]notes.Message, error)
	GetById(userId, messageId int) (notes.Message, error)
	Delete(userId, messageId int) error
	Update(userId, messageId int, input notes.UpdateDeviceMessageInput) error
}

type Repository struct {
	Authorization
	Device
	DevicePhoto
	DeviceVideo
	DeviceAudio
	DeviceMessage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Device:        NewDevicePostgres(db),
		DevicePhoto:   NewDevicePhotoPostgres(db),
		DeviceVideo:   NewDeviceVideoPostgres(db),
		DeviceAudio:   NewDeviceAudioPostgres(db),
		DeviceMessage: NewDeviceMessagePostgres(db),
	}
}
