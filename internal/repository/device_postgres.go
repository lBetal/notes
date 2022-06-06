package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lBetal/notes"
	"github.com/sirupsen/logrus"
)

type DevicePostgres struct {
	db *sqlx.DB
}

func NewDevicePostgres(db *sqlx.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}

func (r *DevicePostgres) Create(userId int, list notes.Device) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (phone_model, phone_number, identification, imei_code) VALUES ($1, $2, $3, $4) RETURNING id", deviceTable)
	row := tx.QueryRow(createListQuery, list.PhoneModel, list.PhoneNumber, list.Indentification, list.ImeiCode)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, device_id) VALUES ($1, $2)", usersDevicesTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		fmt.Println("Tut")
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *DevicePostgres) GetAll(userId int) ([]notes.Device, error) {
	var lists []notes.Device

	query := fmt.Sprintf("SELECT d.id, d.phone_model, d.phone_number, d.identification, d.imei_code FROM %s d INNER JOIN %s ud on d.id = ud.device_id WHERE ud.user_id = $1",
		deviceTable, usersDevicesTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *DevicePostgres) GetById(userId, listId int) (notes.Device, error) {
	var list notes.Device

	query := fmt.Sprintf(`SELECT d.id, d.phone_model, d.phone_number, d.identification, d.imei_code FROM %s d
								INNER JOIN %s ud on d.id = ud.device_id WHERE ud.user_id = $1 AND ud.device_id = $2`,
		deviceTable, usersDevicesTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *DevicePostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s d USING %s ud WHERE d.id = ud.device_id AND ud.user_id=$1 AND ud.device_id=$2",
		deviceTable, usersDevicesTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *DevicePostgres) Update(userId, listId int, input notes.UpdateDeviceInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.PhoneModel != nil {
		setValues = append(setValues, fmt.Sprintf("phone_model=$%d", argId))
		args = append(args, *input.PhoneModel)
		argId++
	}

	if input.PhoneNumber != nil {
		setValues = append(setValues, fmt.Sprintf("phone_number=$%d", argId))
		args = append(args, *input.PhoneNumber)
		argId++
	}

	if input.Indentification != nil {
		setValues = append(setValues, fmt.Sprintf("identification=$%d", argId))
		args = append(args, *input.Indentification)
		argId++
	}

	if input.ImeiCode != nil {
		setValues = append(setValues, fmt.Sprintf("imei_code=$%d", argId))
		args = append(args, *input.ImeiCode)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ud WHERE tl.id = ud.device_id AND ud.device_id=$%d AND ud.user_id=$%d",
		deviceTable, setQuery, usersDevicesTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
