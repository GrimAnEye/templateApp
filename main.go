package main

import (
	c "templateApp/configs"
	"templateApp/logs"

	"go.uber.org/zap"
)

func main() {

	// Создание базового логера
	logs.General()

	// Загрузка настроек из файла, либо создание шаблона
	if err := c.GetConfig(&c.Settings, "configs"); err != nil {
		zap.L().Fatal("ошибка загрузки настроек из файла", zap.Error(err))
	}

	// После получения настроек, обновление логера
	zap.L().Sync()
	logs.Reconfigured()
	defer zap.L().Sync()

	zap.L().Info("Подготовка окончена")
	zap.L().Error("тестовая ошибка")

	// а тут начинается основная деятельность

}
