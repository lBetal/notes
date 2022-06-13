package service

import (
	"github.com/lBetal/notes"
	"github.com/lBetal/notes/internal/repository"
)

type DeviceAudioService struct {
	repo       repository.DeviceAudio
	DeviceRepo repository.Device
}

func NewDeviceAudioService(repo repository.DeviceAudio, DeviceRepo repository.Device) *DeviceAudioService {
	return &DeviceAudioService{repo: repo, DeviceRepo: DeviceRepo}
}

func (s *DeviceAudioService) Create(userId, deviceId int, item notes.Audio) (int, error) {
	_, err := s.DeviceRepo.GetById(userId, deviceId)
	if err != nil {
		// device does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Create(deviceId, item)
}

func (s *DeviceAudioService) GetAll(userId, deviceId int) ([]notes.Audio, error) {
	return s.repo.GetAll(userId, deviceId)
}

func (s *DeviceAudioService) GetById(userId, audioId int) (notes.Audio, error) {
	return s.repo.GetById(userId, audioId)
}

func (s *DeviceAudioService) Delete(userId, audioId int) error {
	return s.repo.Delete(userId, audioId)
}

func (s *DeviceAudioService) Update(userId, audioId int, input notes.UpdateDeviceAudioInput) error {
	return s.repo.Update(userId, audioId, input)
}
