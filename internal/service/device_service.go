package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type DeviceService struct {
	repo repository.Device
}

func NewDeviceService(repo repository.Device) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) Create(userId int, device notes.Device) (int, error) {
	return s.repo.Create(userId, device)
}

func (s *DeviceService) GetAll(userId int) ([]notes.Device, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceService) GetById(userId, deviceId int) (notes.Device, error) {
	return s.repo.GetById(userId, deviceId)
}

func (s *DeviceService) Delete(userId, deviceId int) error {
	return s.repo.Delete(userId, deviceId)
}

func (s *DeviceService) Update(userId, deviceId int, input notes.UpdateDeviceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, deviceId, input)
}
