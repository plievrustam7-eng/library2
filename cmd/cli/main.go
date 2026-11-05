package main

import (
	"fmt"
	"os"

	"library-app/library"
	"library-app/storage"
)

func main() {
	const sourceFile = "books_n_readers.json"

	currentCatalog, err := storage.ReadData(sourceFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("[Система]: Локальный файл не обнаружен. Инициализация пустого каталога.")
			currentCatalog = library.NewCatalog()
		} else {
			fmt.Fprintf(os.Stderr, "Критическая ошибка разбора файла данных: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("[Система]: Данные успешно синхронизированы с базой.")
	}

	ExecuteUI(currentCatalog, sourceFile)
}
