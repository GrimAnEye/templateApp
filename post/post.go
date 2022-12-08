package post

import (
	"fmt"
	"net/smtp"
)

func Send(
	host string,
	port uint16,

	sender,
	recipient,

	subject,
	text string,
) error {
	// Подключаюсь к smtp серверу
	c, err := smtp.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return fmt.Errorf("ошибка подключения к почтовому серверу: %w", err)
	}

	// Назначаю отправителя
	if err := c.Mail(sender); err != nil {
		return fmt.Errorf("ошибка назначения отправителя: %w", err)
	}

	// Назначаю получателя
	if err := c.Rcpt(recipient); err != nil {
		return fmt.Errorf("ошибка назначения получателя: %w", err)
	}
	// Создаю тело сообщения
	wc, err := c.Data()
	if err != nil {
		return fmt.Errorf("ошибка создания тела сообщения: %w", err)
	}

	// Отправляю тело сообщения
	if _, err = fmt.Fprintf(wc,
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n\r\n"+
			"%s",
		sender, recipient, subject, text,
	); err != nil {
		return fmt.Errorf("ошибка отправки данных тела сообщения: %w", err)
	}

	// Закрываю добавление текста в тело письма
	if err := wc.Close(); err != nil {
		return fmt.Errorf("ошибка закрытия подключения к телу сообщения: %w", err)
	}

	// Выходом так же производится отправка письма
	if err = c.Quit(); err != nil {
		return fmt.Errorf("ошибка при окончании работы с почтовым сервером: %w", err)
	}
	return nil

}
