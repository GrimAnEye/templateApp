package configs

import (
	"errors"
	"fmt"
	"os"
	p "path/filepath"
	"reflect"

	yaml "gopkg.in/yaml.v3"
)

// GetConfig - получает указатель на ИНИЦИАЛИЗИРОВАННУЮ переменную конфигурации и путь,
// по которому она должна располагаться в файловой системе, относительно родительской директории.
//
// Проверяет наличие файла конфигурации по указанному пути,
//   - при его отсутствии - создает каталог и файл конфигурации в формате <TypeName>.yaml.example, содержащий экспортируемые поля.
//     Каталог создается в правами 0744, файл - 0640
//   - при наличии - открывает файл и демаршализирует данные в указанную переменную.
//     Нет гарантии, что имеющиеся в переменной данные, не будут перезаписаны в процессе разбора файла конфигурации
func GetConfig[T any](v *T, configPath string) error {

	var name string = reflect.TypeOf(*v).Name()

	// Создание каталога настроек
	if err := os.MkdirAll(p.FromSlash(configPath), 0744); err != nil {
		return fmt.Errorf("ошибка создания каталога настроек: %w", err)
	}

	// Проверка существования файла настроек
	if _, err := os.Stat(p.Join(configPath, name+".yaml")); errors.Is(err, os.ErrNotExist) {

		// В случае отсутствия файла настроек - создаю пустую структуру и кладу её в файл
		return newConfig(v, &name, &configPath)

		// В случае иных ошибок - возвращаю их
	} else if err != nil {
		return fmt.Errorf("ошибка проверки существования файла настроек: %w", err)
	}

	// Если файл настроек есть - начинаю загрузку параметров
	return loadConfig(v, &name, &configPath)
}

func newConfig[T any](v *T, name, configPath *string) error {

	// Формирую пустую структуру данных
	data, err := yaml.Marshal(v)
	if err != nil {

		return fmt.Errorf("ошибка маршалинга полученной структуры %s: %w", *name, err)
	}

	// Записываю шаблон в файл
	if err = os.WriteFile(p.Join(*configPath, *name+".yaml.example"), data, 0640); err != nil {

		return fmt.Errorf("ошибка записи структуры в файл %s.yaml.example: %w", *name, err)
	}
	return fmt.Errorf("создан файл шаблона %s.yaml.example", *name)
}

func loadConfig[T any](v *T, name, configPath *string) error {

	// Получаю объект *.yaml
	co, err := os.Open(p.Join(*configPath, *name+".yaml"))
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s.yaml: %w", *name, err)
	}
	defer co.Close()

	// Разбираю данные из файла
	decoder := yaml.NewDecoder(co)
	if err := decoder.Decode(v); err != nil {
		return fmt.Errorf("ошибка разбора файла %s.yaml: %w", *name, err)
	}

	return nil
}
