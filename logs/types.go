package logs

type TLogs struct {
	AlterLogPathFolder string // Путь до файла журнала. Если пустой, то журнал складывается в каталог программы
	MaxSizeMb          int    // Вес одного файла журнала
	MaxBackupsCount    int    // Число файлов для сохранения
	MaxAgeDays         int    // Число дней, до удаления файла
	Debug              bool   // Вывод отладочной информации в журнал
}
