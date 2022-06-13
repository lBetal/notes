package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type DeviceMessageService struct {
	repo       repository.DeviceMessage
	DeviceRepo repository.Device
}

func NewDeviceMessageService(repo repository.DeviceMessage, DeviceRepo repository.Device) *DeviceMessageService {
	return &DeviceMessageService{repo: repo, DeviceRepo: DeviceRepo}
}

func (s *DeviceMessageService) Create(userId, deviceId int, item notes.Message) (int, error) {
	_, err := s.DeviceRepo.GetById(userId, deviceId)
	if err != nil {
		// device does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(deviceId, item)
}

func (s *DeviceMessageService) GetAll(userId, deviceId int) ([]notes.Message, error) {
	return s.repo.GetAll(userId, deviceId)
}

func (s *DeviceMessageService) GetById(userId, messageId int) (notes.Message, error) {
	return s.repo.GetById(userId, messageId)
}

func (s *DeviceMessageService) Delete(userId, messageId int) error {
	return s.repo.Delete(userId, messageId)
}

func (s *DeviceMessageService) Update(userId, messageId int, input notes.UpdateDeviceMessageInput) error {
	return s.repo.Update(userId, messageId, input)
}
