package foldercash

import (
	unqCash "bring_some_water_please/utils/UniqueCash"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const cashFolder string = "internal/data/cashfile/"

func NewCashFolder(userId int64) (folder string, err error) {
	seconds := unqCash.NewCash(userId)
	folder = cashFolder + fmt.Sprint(seconds)

	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return "", err
	}

	return folder, nil
}

func DeleteCashFolder(folder string) error {
	err := os.RemoveAll(folder)

	if err != nil {
		log.Printf("Ошибка при удалении папки и её содержимого: %v", err)
		return err
	}

	log.Printf("Папка %s и всё её содержимое удалено", folder)
	return nil
}

func AddFileInCashFolder(data []byte, folder, fileName string) error {
	filePath := filepath.Join(folder, fileName+".jar")

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении файла %s: %w", filePath, err)
	}

	log.Printf("Файл успешно сохранён: %s", filePath)
	return nil
}
