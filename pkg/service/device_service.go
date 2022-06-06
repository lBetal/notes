package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/pkg/repository"
)

type DeviceService struct {
	repo repository.Device
}

func NewDeviceService(repo repository.Device) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) Create(userId int, list notes.Device) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *DeviceService) GetAll(userId int) ([]notes.Device, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceService) GetById(userId, listId int) (notes.Device, error) {
	return s.repo.GetById(userId, listId)
}

func (s *DeviceService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *DeviceService) Update(userId, listId int, input notes.UpdateDeviceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
