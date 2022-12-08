package configs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	yaml "gopkg.in/yaml.v3"
)

// LoadConfig - проверяет наличие файла настроек из файла, либо создаёт пустой шаблон
func LoadConfig[T any](v T) (*T, error) {

	// Создание каталога настроек
	if err := os.MkdirAll(filepath.FromSlash(PathConfigs), 0744); err != nil {
		return nil, fmt.Errorf("ошибка создания каталога настроек: %w", err)
	}

	// Проверка существования файла настроек
	if _, err := os.Stat(
		filepath.FromSlash(
			PathConfigs + "/" + reflect.TypeOf(v).Name() + ".yaml",
		),
	); errors.Is(err, os.ErrNotExist) {

		// В случае отсутствия файла настроек - создаю пустую структуру и кладу её в файл
		return newConfig(v)

		// В случае иных ошибок - возвращаю их
	} else if err != nil {
		return nil, fmt.Errorf("ошибка проверки существования файла настроек: %w", err)
	}

	// Если файл настроек есть - начинаю загрузку параметров
	return loadConfig(v)
}

func newConfig[T any](v T) (*T, error) {

	// Формирую пустую структуру данных
	data, err := yaml.Marshal(&v)
	if err != nil {

		return nil, fmt.Errorf("ошибка маршалинга полученной структуры %T: %w", v, err)
	}

	// Записываю шаблон в файл
	if err = os.WriteFile(
		filepath.FromSlash(PathConfigs+"/"+reflect.TypeOf(v).Name()+".yaml.example"),
		data,
		0640); err != nil {

		return nil, fmt.Errorf("ошибка записи структуры в файл %T.yaml.example: %w", v, err)
	}

	return nil, fmt.Errorf("создан файл шаблона %T.yaml.example", v)
}

func loadConfig[T any](v T) (*T, error) {

	// Получаю объект *.yaml
	co, err := os.Open(PathConfigs + "/" + reflect.TypeOf(v).Name() + ".yaml")
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла %T.yaml: %w", v, err)
	}
	defer co.Close()

	// Разбираю данные из файла
	var s T
	decoder := yaml.NewDecoder(co)
	if err := decoder.Decode(&s); err != nil {
		return nil, fmt.Errorf("ошибка разбора файла %T.yaml: %w", v, err)
	}

	return &s, nil
}
