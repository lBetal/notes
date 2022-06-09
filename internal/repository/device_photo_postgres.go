package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lBetal/notes"
)

type DevicePhotoPostgres struct {
	db *sqlx.DB
}

func NewDevicePhotoPostgres(db *sqlx.DB) *DevicePhotoPostgres {
	return &DevicePhotoPostgres{db: db}
}

func (r *DevicePhotoPostgres) Create(deviceId int, item notes.Photo) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {

		return 0, err
	}

	var PhotoId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (path) values ($1) RETURNING id", photoTable)

	row := tx.QueryRow(createItemQuery, item.Path)
	err = row.Scan(&PhotoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (device_id, photo_id) values ($1, $2)", devicesPhotosTable)
	_, err = tx.Exec(createListItemsQuery, deviceId, PhotoId)
	if err != nil {

		tx.Rollback()
		return 0, err
	}

	return PhotoId, tx.Commit()
}

func (r *DevicePhotoPostgres) GetAll(userId, deviceId int) ([]notes.Photo, error) {
	var photos []notes.Photo
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.photo_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE d.device_id = $1 AND ul.user_id = $2`,
		photoTable, devicesPhotosTable, usersDevicesTable)
	if err := r.db.Select(&photos, query, deviceId, userId); err != nil {
		return nil, err
	}

	return photos, nil
}

func (r *DevicePhotoPostgres) GetById(userId, photoId int) (notes.Photo, error) {
	var photo notes.Photo
	query := fmt.Sprintf(`SELECT p.id, p.path FROM %s p INNER JOIN %s d on d.photo_id = p.id
									INNER JOIN %s ul on ul.device_id = d.device_id WHERE p.id = $1 AND ul.user_id = $2`,
		photoTable, devicesPhotosTable, usersDevicesTable)
	if err := r.db.Get(&photo, query, photoId, userId); err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *DevicePhotoPostgres) Delete(userId, photoId int) error {
	query := fmt.Sprintf(`DELETE FROM %s p USING %s d, %s ul 
									WHERE p.id = d.photo_id AND d.device_id = ul.device_id AND ul.user_id = $1 AND p.id = $2`,
		photoTable, devicesPhotosTable, usersDevicesTable)
	_, err := r.db.Exec(query, userId, photoId)
	return err
}

func (r *DevicePhotoPostgres) Update(userId, photoId int, input notes.UpdateDevicePhotoInput) error {
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
									WHERE p.id = d.photo_id AND d.device_id = ul.device_id AND ul.user_id = $%d AND p.id = $%d`,
		photoTable, setQuery, devicesPhotosTable, usersDevicesTable, argId, argId+1)
	args = append(args, userId, photoId)

	_, err := r.db.Exec(query, args...)
	return err
}
