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
	Create(userId int, deivce notes.Device) (int, error)
	GetAll(userId int) ([]notes.Device, error)
	GetById(userId, deivceId int) (notes.Device, error)
	Delete(userId, deivceId int) error
	Update(userId, deivceId int, input notes.UpdateDeviceInput) error
}

type DevicePhoto interface {
	Create(userId, deviceId int, item notes.Photo) (int, error)
	GetAll(userId, deviceId int) ([]notes.Photo, error)
	GetById(userId, photoId int) (notes.Photo, error)
	Delete(userId, photoId int) error
	Update(userId, photoId int, input notes.UpdateDevicePhotoInput) error
}

type DeviceVideo interface {
	Create(userId, deviceId int, item notes.Video) (int, error)
	GetAll(userId, deviceId int) ([]notes.Video, error)
	GetById(userId, videoId int) (notes.Video, error)
	Delete(userId, videoId int) error
	Update(userId, videoId int, input notes.UpdateDeviceVideoInput) error
}

type DeviceAudio interface {
	Create(userId, deviceId int, item notes.Audio) (int, error)
	GetAll(userId, deviceId int) ([]notes.Audio, error)
	GetById(userId, audioId int) (notes.Audio, error)
	Delete(userId, audioId int) error
	Update(userId, audioId int, input notes.UpdateDeviceAudioInput) error
}

type DeviceMessage interface {
	Create(userId, deviceId int, item notes.Message) (int, error)
	GetAll(userId, deviceId int) ([]notes.Message, error)
	GetById(userId, messageId int) (notes.Message, error)
	Delete(userId, messageId int) error
	Update(userId, messageId int, input notes.UpdateDeviceMessageInput) error
}

type Service struct {
	Authorization
	Device
	DevicePhoto
	DeviceVideo
	DeviceAudio
	DeviceMessage
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Device:        NewDeviceService(repos.Device),
		DevicePhoto:   NewDevicePhotoService(repos.DevicePhoto, repos.Device),
		DeviceVideo:   NewDeviceVideoService(repos.DeviceVideo, repos.Device),
		DeviceAudio:   NewDeviceAudioService(repos.DeviceAudio, repos.Device),
		DeviceMessage: NewDeviceMessageService(repos.DeviceMessage, repos.Device),
	}
}
