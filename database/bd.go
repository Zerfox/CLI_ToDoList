package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// Генерит 2 файла под хранение тудушек и событий
func WorkList(dir string) {

	DataList, err := os.Stat(dir + "dataList.txt")
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не существует
			fmt.Println("Файл отсутствует, будет создан новый файл dataList")
			dataList, err := os.Create(dir + "dataList.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				if err := dataList.Close(); err != nil {
					log.Printf("Ошибка при закрытии файла dataList: %v", err)
				}
			}()
		} else {
			// Другая ошибка при попытке получить информацию о файле
			fmt.Println("Ошибка при проверке файла задач:", err)
		}
	} else {
		// Файл существует, info содержит метаданные файла
		fmt.Println("Файл подключен, его размер:", DataList.Size())
	}

	DataEvent, err := os.Stat(dir + "dataEvent.txt")
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не существует
			fmt.Println("Файл отсутствует, будет создан новый файл dataEvent")
			dataEvent, err := os.Create(dir + "dataEvent.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer func() {
				if err := dataEvent.Close(); err != nil {
					log.Printf("Ошибка при закрытии файла DataEvent: %v", err)
				}
			}()
		} else {
			// Другая ошибка при попытке получить информацию о файле
			fmt.Println("Ошибка при проверке файла событий:", err)
		}
	} else {
		// Файл существует, info содержит метаданные файла
		fmt.Println("Файл подключен, его размер:", DataEvent.Size())
	}

}

func Save(path string, v any) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0644)
}

func Load[T any](path string, out *T) error {
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil // первого файла может не быть — это ок
		}
		return err
	}
	if len(b) == 0 { // пустой файл — просто выходим
		return nil
	}
	return json.Unmarshal(b, out)
}
