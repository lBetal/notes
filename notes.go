package notes

import "errors"

type Device struct {
	Id              int    `json:"id" db:"id"`
	PhoneModel      string `json:"phone_model" db:"phone_model" binding:"required"`
	PhoneNumber     uint64 `json:"phone_number" db:"phone_number"`
	Indentification uint64 `json:"identification" db:"identification"`
	ImeiCode        uint64 `json:"imei_code" db:"imei_code"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type DeviceItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type DeviceListsItem struct {
	Id     int
	ListId int
	ItemId int
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

type UpdateDeviceItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateDeviceItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
