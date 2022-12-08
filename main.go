package main

import (
	c "templateApp/configs"
	"templateApp/logs"

	"go.uber.org/zap"
)

func main() {

	// Создание базового логера
	logs.General()
	//defer zap.L().Sync()

	// Загрузка настроек из файла, либо создание шаблона
	if settings, err := c.LoadConfig(c.TSettings{}); err != nil {
		zap.L().Fatal("ошибка загрузки настроек из файла", zap.Error(err))
	} else {
		c.Settings = settings
	}

	// После получения настроек, обновление логера
	zap.L().Sync()
	logs.Reconfigured()
	defer zap.L().Sync()

	zap.L().Info("Подготовка окончена")
	zap.L().Error("тестовая ошибка")

	// а тут начинается основная деятельность

}
