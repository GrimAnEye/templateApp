package main

import (
	c "templateApp/configs"
	l "templateApp/logs"
	d "templateApp/modules/database"
	m "templateApp/modules/mail"

	"go.uber.org/zap"
)

func main() {

	// Создание базового логера
	l.General()

	// Загрузка настроек из файла, либо создание шаблона
	if err := c.GetConfig(&l.LogsConf, "configs"); err != nil {
		zap.L().Fatal("ошибка загрузки настроек логгера из файла", zap.Error(err))
	}
	if err := c.GetConfig(&m.MailConf, "configs"); err != nil {
		zap.L().Fatal("ошибка загрузки настроек логгера из файла", zap.Error(err))
	}
	if err := c.GetConfig(&d.DatabaseConf, "configs"); err != nil {
		zap.L().Fatal("ошибка загрузки настроек логгера из файла", zap.Error(err))
	}

	// После получения настроек, обновление логера
	zap.L().Sync()
	l.Reconfigured(m.MailConf.Host, m.MailConf.Port, m.MailConf.ErrorsSender, m.MailConf.ErrorsRecipient, m.MailConf.ErrorsSubject)
	defer zap.L().Sync()

	zap.L().Info("Подготовка окончена")

	// а тут начинается основная деятельность

}
