package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lBetal/notes"
)

type DeviceVideoPostgres struct {
	db *sqlx.DB
}

func NewDeviceVideoPostgres(db *sqlx.DB) *DeviceVideoPostgres {
	return &DeviceVideoPostgres{db: db}
}

func (r *DeviceVideoPostgres) Create(deviceId int, item notes.Video) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {

		return 0, err
	}

	var VideoId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (path) values ($1) RETURNING id", videoTable)

	row := tx.QueryRow(createItemQuery, item.Path)
	err = row.Scan(&VideoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (device_id, video_id) values ($1, $2)", devicesVideosTable)
	_, err = tx.Exec(createListItemsQuery, deviceId, VideoId)
	if err != nil {

		tx.Rollback()
		return 0, err
	}

	return VideoId, tx.Commit()
}

func (r *DeviceVideoPostgres) GetAll(userId, deviceId int) ([]notes.Video, error) {
	var videos []notes.Video
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.video_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE d.device_id = $1 AND ul.user_id = $2`,
		videoTable, devicesVideosTable, usersDevicesTable)
	if err := r.db.Select(&videos, query, deviceId, userId); err != nil {
		return nil, err
	}

	return videos, nil
}

func (r *DeviceVideoPostgres) GetById(userId, videoId int) (notes.Video, error) {
	var video notes.Video
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.video_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE p.id = $1 AND ul.user_id = $2`,
		videoTable, devicesVideosTable, usersDevicesTable)
	if err := r.db.Get(&video, query, videoId, userId); err != nil {
		return video, err
	}

	return video, nil
}

func (r *DeviceVideoPostgres) Delete(userId, videoId int) error {
	query := fmt.Sprintf(`DELETE FROM %s p USING %s d, %s ul 
									WHERE p.id = d.video_id AND d.device_id = ul.device_id AND ul.user_id = $1 AND p.id = $2`,
		videoTable, devicesVideosTable, usersDevicesTable)
	_, err := r.db.Exec(query, userId, videoId)
	return err
}

func (r *DeviceVideoPostgres) Update(userId, videoId int, input notes.UpdateDeviceVideoInput) error {
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
									WHERE p.id = d.video_id AND d.device_id = ul.device_id AND ul.user_id = $%d AND p.id = $%d`,
		videoTable, setQuery, devicesVideosTable, usersDevicesTable, argId, argId+1)
	args = append(args, userId, videoId)

	_, err := r.db.Exec(query, args...)
	return err
}
