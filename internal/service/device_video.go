package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type DeviceVideoService struct {
	repo       repository.DeviceVideo
	DeviceRepo repository.Device
}

func NewDeviceVideoService(repo repository.DeviceVideo, DeviceRepo repository.Device) *DeviceVideoService {
	return &DeviceVideoService{repo: repo, DeviceRepo: DeviceRepo}
}

func (s *DeviceVideoService) Create(userId, deviceId int, item notes.Video) (int, error) {
	_, err := s.DeviceRepo.GetById(userId, deviceId)
	if err != nil {
		// device does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(deviceId, item)
}

func (s *DeviceVideoService) GetAll(userId, deviceId int) ([]notes.Video, error) {
	return s.repo.GetAll(userId, deviceId)
}

func (s *DeviceVideoService) GetById(userId, videoId int) (notes.Video, error) {
	return s.repo.GetById(userId, videoId)
}

func (s *DeviceVideoService) Delete(userId, videoId int) error {
	return s.repo.Delete(userId, videoId)
}

func (s *DeviceVideoService) Update(userId, videoId int, input notes.UpdateDeviceVideoInput) error {
	return s.repo.Update(userId, videoId, input)
}
