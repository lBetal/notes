package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lBetal/notes"
)

type DeviceMessagePostgres struct {
	db *sqlx.DB
}

func NewDeviceMessagePostgres(db *sqlx.DB) *DeviceMessagePostgres {
	return &DeviceMessagePostgres{db: db}
}

func (r *DeviceMessagePostgres) Create(deviceId int, item notes.Message) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {

		return 0, err
	}

	var messageId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (path) values ($1) RETURNING id", messageTable)

	row := tx.QueryRow(createItemQuery, item.Path)
	err = row.Scan(&messageId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (device_id, message_id) values ($1, $2)", devicesMessagesTable)
	_, err = tx.Exec(createListItemsQuery, deviceId, messageId)
	if err != nil {

		tx.Rollback()
		return 0, err
	}

	return messageId, tx.Commit()
}

func (r *DeviceMessagePostgres) GetAll(userId, deviceId int) ([]notes.Message, error) {
	var messages []notes.Message
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.message_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE d.device_id = $1 AND ul.user_id = $2`,
		messageTable, devicesMessagesTable, usersDevicesTable)
	if err := r.db.Select(&messages, query, deviceId, userId); err != nil {
		return nil, err
	}

	return messages, nil
}

func (r *DeviceMessagePostgres) GetById(userId, messageId int) (notes.Message, error) {
	var message notes.Message
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.message_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE p.id = $1 AND ul.user_id = $2`,
		messageTable, devicesMessagesTable, usersDevicesTable)
	if err := r.db.Get(&message, query, messageId, userId); err != nil {
		return message, err
	}

	return message, nil
}

func (r *DeviceMessagePostgres) Delete(userId, messageId int) error {
	query := fmt.Sprintf(`DELETE FROM %s p USING %s d, %s ul 
									WHERE p.id = d.message_id AND d.device_id = ul.device_id AND ul.user_id = $1 AND p.id = $2`,
		messageTable, devicesMessagesTable, usersDevicesTable)
	_, err := r.db.Exec(query, userId, messageId)
	return err
}

func (r *DeviceMessagePostgres) Update(userId, messageId int, input notes.UpdateDeviceMessageInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Path != nil {
		setValues = append(setValues, fmt.Sprintf("path=$%d", argId))
		args = append(args, *input.Path)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s p SET %s FROM %s d, %s ul
									WHERE p.id = d.message_id AND d.device_id = ul.device_id AND ul.user_id = $%d AND p.id = $%d`,
		messageTable, setQuery, devicesMessagesTable, usersDevicesTable, argId, argId+1)
	args = append(args, userId, messageId)

	_, err := r.db.Exec(query, args...)
	return err
}
