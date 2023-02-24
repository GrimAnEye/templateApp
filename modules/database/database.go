// Модуль используется для управления базой данных на базе GORM
package database

import (
	"fmt"
	"time"

	z "go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// Connect - функция подключения к БД, сохраняя объект
// подключения в глобальной переменной
func Connect(
	dbHost string, dbPort uint16,
	dbUser, dbPassword, dbName string,

) (*gorm.DB, error) {

	// Формирование строки подключения
	dsn := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Перенаправление журнала Gorm в Zap
	logger := zapgorm2.New(z.L())
	logger.SetAsDefault()

	// Отрытие подключения в БД
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {

		z.L().Error("ошибка подключения к БД",
			z.String("host", dbHost),
			z.Uint16("port", dbPort),
			z.String("dbName", dbName),
			z.String("user", dbUser),
			z.Error(err))
		return nil, err
	}

	// Получение сырого объекта к БД, для настройки пула подключений
	rawDb, err := db.DB()
	if err != nil {
		z.L().Error("ошибка при получении сырого подключения БД", z.Error(err))
		return nil, err
	}
	rawDb.SetMaxIdleConns(2)
	rawDb.SetConnMaxIdleTime(time.Minute * 30)

	rawDb.SetMaxOpenConns(20)
	rawDb.SetConnMaxLifetime(time.Hour)

	return db, nil
}
