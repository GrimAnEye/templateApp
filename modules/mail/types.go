package mail

type TMail struct {
	Host            string // Адрес хоста, smpt.example.com
	Port            uint16 // Порт подключения
	ErrorsSender    string // От какого адреса будет отправляться почта
	ErrorsRecipient string // Куда будет отправляться почта
	ErrorsSubject   string // Тема письма
}
