package service

import (
	"github.com/lBetal/todo"
	"github.com/lBetal/todo/pkg/repository"
)

type DeviceItemService struct {
	repo       repository.DeviceItem
	DeviceRepo repository.Device
}

func NewDeviceItemService(repo repository.DeviceItem, DeviceRepo repository.Device) *DeviceItemService {
	return &DeviceItemService{repo: repo, DeviceRepo: DeviceRepo}
}

func (s *DeviceItemService) Create(userId, listId int, item todo.DeviceItem) (int, error) {
	_, err := s.DeviceRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *DeviceItemService) GetAll(userId, listId int) ([]todo.DeviceItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *DeviceItemService) GetById(userId, itemId int) (todo.DeviceItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *DeviceItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *DeviceItemService) Update(userId, itemId int, input todo.UpdateDeviceItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
