package notes

import "errors"

type Device struct {
	Id              int    `json:"id" db:"id"`
	PhoneModel      string `json:"phone_model" db:"phone_model" binding:"required"`
	PhoneNumber     uint64 `json:"phone_number" db:"phone_number"`
	Indentification uint64 `json:"identification" db:"identification"`
	ImeiCode        uint64 `json:"imei_code" db:"imei_code"`
}

type UsersDevice struct {
	Id       int
	UserId   int
	DeviceId int
}

type Photo struct {
	Id   int    `json:"id" db:"id"`
	Path string `json:"path" db:"path" binding:"required"`
}

type DevicePhoto struct {
	Id       int
	DeviceId int
	PhotoId  int
}

type UpdateDeviceInput struct {
	PhoneModel      *string `json:"phone_model" db:"phone_model" binding:"required"`
	PhoneNumber     *uint64 `json:"phone_number" db:"phone_number"`
	Indentification *uint64 `json:"identification" db:"identification"`
	ImeiCode        *uint64 `json:"imei_code" db:"imei_code"`
}

func (u UpdateDeviceInput) Validate() error {
	if u.PhoneModel == nil && u.PhoneNumber == nil && u.Indentification == nil && u.ImeiCode == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateDevicePhotoInput struct {
	Path *string `json:"path"`
}

func (i UpdateDevicePhotoInput) Validate() error {
	if i.Path == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
