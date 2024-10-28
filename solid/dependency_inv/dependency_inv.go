package main

import "fmt"

// Интерфейс для абстракции хранения данных
type DataStorage interface {
	Save(data string)
}

// Низкоуровневый модуль для хранения данных в файле
type FileStorage struct{}

func (fs *FileStorage) Save(data string) {
	fmt.Println("Сохранение данных в файл:", data)
}

// Высокоуровневый модуль, не зависит напрямую от FileStorage
type DataManager struct {
	storage DataStorage // зависит от абстракции
}

func (dm *DataManager) SaveData(data string) {
	dm.storage.Save(data) // делегирование сохранения
}

func main() {
	fs := &FileStorage{}
	dm := DataManager{storage: fs}
	dm.SaveData("Тестовые данные")
}
