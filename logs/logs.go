package logs

import (
	"fmt"
	c "templateApp/configs"
	"templateApp/mail"

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
//	отправляет уведомления на почту
//	проводит ротацию устаревших журналов
func Reconfigured() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/logs.log",
		MaxSize:    c.Settings.Logs.MaxSize,
		MaxBackups: c.Settings.Logs.MaxBackups,
		MaxAge:     c.Settings.Logs.MaxAgeDays,
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), w,
		func() zapcore.LevelEnabler {
			if c.Settings.Logs.Debug {
				return zap.DebugLevel
			}
			return zap.InfoLevel
		}(),
	)

	zap.ReplaceGlobals(
		zap.New(
			core,
			zap.AddStacktrace(zap.ErrorLevel),
			zap.Hooks(func(e zapcore.Entry) error {
				if e.Level >= zap.ErrorLevel {
					return mail.Send(
						c.Settings.Mail.Host,
						c.Settings.Mail.Port,
						c.Settings.Mail.Sender,
						c.Settings.Mail.Recipient,
						c.Settings.Mail.Subject,
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
				return nil
			}),
		),
	)

}
