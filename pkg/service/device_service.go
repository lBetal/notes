package service

import (
	"github.com/lBetal/todo"
	"github.com/lBetal/todo/pkg/repository"
)

type DeviceService struct {
	repo repository.Device
}

func NewDeviceService(repo repository.Device) *DeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) Create(userId int, list todo.Device) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *DeviceService) GetAll(userId int) ([]todo.Device, error) {
	return s.repo.GetAll(userId)
}

func (s *DeviceService) GetById(userId, listId int) (todo.Device, error) {
	return s.repo.GetById(userId, listId)
}

func (s *DeviceService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *DeviceService) Update(userId, listId int, input todo.UpdateDeviceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, listId, input)
}
