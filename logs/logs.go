package logs

import (
	"fmt"
	"path/filepath"
	"templateApp/modules/mail"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// General - элементарный логгер, используемый сразу после запуска
func General() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

// Reconfigured - логгер, использующий параметры из файла настроек:
//
//	ведёт файлы журнала
//	отправляет уведомления на почту, в случае появления ошибок
//	проводит ротацию устаревших журналов
func Reconfigured(

	mailHost string, mailPort uint16,
	mailSender, mailRecipient, mailSubject string,
) {

	// Добавляю ротацию журналов
	w := zapcore.AddSync(&lumberjack.Logger{

		Filename: filepath.Join(func() string {
			if LogsConf.AlterLogPathFolder != "" {
				return LogsConf.AlterLogPathFolder
			}
			return LogPathDefault
		}(), "logs.log"),

		MaxSize:    LogsConf.MaxSizeMb,
		MaxBackups: LogsConf.MaxBackupsCount,
		MaxAge:     LogsConf.MaxAgeDays,
	})

	// Определяю уроверь подробности журнала
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), w,
		func() zapcore.LevelEnabler {
			if LogsConf.Debug {
				return zap.DebugLevel
			}
			return zap.InfoLevel
		}(),
	)

	// Формирую структуру логгера
	zap.ReplaceGlobals(zap.New(core, zap.AddStacktrace(zap.ErrorLevel),

		zap.Hooks(func(e zapcore.Entry) error {

			// Если задан почтовый хост для отправки уведомлений
			// и сообщение об ошибке или хуже, отправляю уведомление на почту
			if mailHost != "" {

				if e.Level >= zap.ErrorLevel {
					return mail.Send(
						mailHost, mailPort,
						mailSender, mailRecipient, mailSubject,

						fmt.Sprintf("Error found:\r\n\r\n"+
							"Level: %d\r\n"+
							"Time: %s\r\n"+
							"LoggerName: %s\r\n"+
							"Message: %s\r\n"+
							"Caller: %s\r\n"+
							"Stack: %s\r\n",
							e.Level,
							e.Time.Format("2006.01.02 15:04:05 -07:00"),
							e.LoggerName,
							e.Message,
							e.Caller.FullPath(),
							e.Stack))
				}
			}
			return nil
		}),
	),
	)
}
