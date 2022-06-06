package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type DeviceItemService struct {
	repo       repository.DeviceItem
	DeviceRepo repository.Device
}

func NewDeviceItemService(repo repository.DeviceItem, DeviceRepo repository.Device) *DeviceItemService {
	return &DeviceItemService{repo: repo, DeviceRepo: DeviceRepo}
}

func (s *DeviceItemService) Create(userId, listId int, item notes.DeviceItem) (int, error) {
	_, err := s.DeviceRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *DeviceItemService) GetAll(userId, listId int) ([]notes.DeviceItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *DeviceItemService) GetById(userId, itemId int) (notes.DeviceItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *DeviceItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *DeviceItemService) Update(userId, itemId int, input notes.UpdateDeviceItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
