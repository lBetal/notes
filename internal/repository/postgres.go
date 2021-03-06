package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable           = "users"
	deviceTable          = "device"
	usersDevicesTable    = "users_devices"
	photoTable           = "photo"
	devicesPhotosTable   = "devices_photos"
	videoTable           = "video"
	devicesVideosTable   = "devices_videos"
	audioTable           = "audio"
	devicesAudiosTable   = "devices_audios"
	messageTable         = "message"
	devicesMessagesTable = "devices_messages"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
