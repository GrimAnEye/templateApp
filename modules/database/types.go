package database

import "time"

// Database - настройки подключения к БД
type TDatabase struct {
	User     string // Имя пользователя при подключении к БД
	Password string // Пароль пользователя при подключении к БД
	Host     string // Имя или IP хост-машины для подключения к БД
	Port     uint16 // Порт для подключения к БД
	Name     string // Имя базы для подключения

	MaxIdleConns    int           // https://pkg.go.dev/database/sql#DB.SetMaxIdleConns
	ConnMaxIdleTime time.Duration // https://pkg.go.dev/database/sql#DB.SetConnMaxIdleTime

	MaxOpenConns    int           // https://pkg.go.dev/database/sql#DB.SetMaxOpenConns
	ConnMaxLifeTime time.Duration // https://pkg.go.dev/database/sql#DB.SetConnMaxLifetime
}
