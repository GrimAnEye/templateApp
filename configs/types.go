package configs

type TSettings struct {
	Logs struct {
		MaxSize    int  // Вес одного файла журнала
		MaxBackups int  // Число файлов для сохранения
		MaxAgeDays int  // Число дней, до удаления файла
		Debug      bool // Вывод отладочной информации в журнал
	}

	Post struct {
		Host      string // Адрес хоста, smpt.example.com
		Port      uint16 // Порт подключения
		Sender    string // От какого адреса будет отправляться почта
		Recipient string // Куда будет отправляться почта
		Subject   string // Тема письма
	}
}
