package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type DevicePhotoService struct {
	repo       repository.DevicePhoto
	DeviceRepo repository.Device
}

func NewDevicePhotoService(repo repository.DevicePhoto, DeviceRepo repository.Device) *DevicePhotoService {
	return &DevicePhotoService{repo: repo, DeviceRepo: DeviceRepo}
}

func (s *DevicePhotoService) Create(userId, deviceId int, item notes.Photo) (int, error) {
	_, err := s.DeviceRepo.GetById(userId, deviceId)
	if err != nil {
		// device does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(deviceId, item)
}

func (s *DevicePhotoService) GetAll(userId, deviceId int) ([]notes.Photo, error) {
	return s.repo.GetAll(userId, deviceId)
}

func (s *DevicePhotoService) GetById(userId, photoId int) (notes.Photo, error) {
	return s.repo.GetById(userId, photoId)
}

func (s *DevicePhotoService) Delete(userId, photoId int) error {
	return s.repo.Delete(userId, photoId)
}

func (s *DevicePhotoService) Update(userId, photoId int, input notes.UpdateDevicePhotoInput) error {
	return s.repo.Update(userId, photoId, input)
}
