package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lBetal/notes"
)

type DeviceAudioPostgres struct {
	db *sqlx.DB
}

func NewDeviceAudioPostgres(db *sqlx.DB) *DeviceAudioPostgres {
	return &DeviceAudioPostgres{db: db}
}

func (r *DeviceAudioPostgres) Create(deviceId int, item notes.Audio) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {

		return 0, err
	}

	var audioId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (path) values ($1) RETURNING id", audioTable)

	row := tx.QueryRow(createItemQuery, item.Path)
	err = row.Scan(&audioId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (device_id, audio_id) values ($1, $2)", devicesAudiosTable)
	_, err = tx.Exec(createListItemsQuery, deviceId, audioId)
	if err != nil {

		tx.Rollback()
		return 0, err
	}

	return audioId, tx.Commit()
}

func (r *DeviceAudioPostgres) GetAll(userId, deviceId int) ([]notes.Audio, error) {
	var audios []notes.Audio
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.audio_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE d.device_id = $1 AND ul.user_id = $2`,
		audioTable, devicesAudiosTable, usersDevicesTable)
	if err := r.db.Select(&audios, query, deviceId, userId); err != nil {
		return nil, err
	}

	return audios, nil
}

func (r *DeviceAudioPostgres) GetById(userId, audioId int) (notes.Audio, error) {
	var audio notes.Audio
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.audio_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE p.id = $1 AND ul.user_id = $2`,
		audioTable, devicesAudiosTable, usersDevicesTable)
	if err := r.db.Get(&audio, query, audioId, userId); err != nil {
		return audio, err
	}

	return audio, nil
}

func (r *DeviceAudioPostgres) Delete(userId, audioId int) error {
	query := fmt.Sprintf(`DELETE FROM %s p USING %s d, %s ul 
									WHERE p.id = d.audio_id AND d.device_id = ul.device_id AND ul.user_id = $1 AND p.id = $2`,
		audioTable, devicesAudiosTable, usersDevicesTable)
	_, err := r.db.Exec(query, userId, audioId)
	return err
}

func (r *DeviceAudioPostgres) Update(userId, audioId int, input notes.UpdateDeviceAudioInput) error {
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
									WHERE p.id = d.audio_id AND d.device_id = ul.device_id AND ul.user_id = $%d AND p.id = $%d`,
		audioTable, setQuery, devicesAudiosTable, usersDevicesTable, argId, argId+1)
	args = append(args, userId, audioId)

	_, err := r.db.Exec(query, args...)
	return err
}
