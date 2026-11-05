package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"library-app/domain"
	"library-app/library"
)

type FileModel struct {
	Records []*domain.Publication `json:"records_list"`
	Users   []*domain.Client      `json:"users_list"`
}

func WriteData(targetPath string, cat *library.Catalog) error {
	payload := FileModel{
		Records: cat.Items,
		Users:   cat.Members,
	}

	// Использование 4 пробелов вместо 2 для уникализации структуры файла
	bytes, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to encode: %w", err)
	}

	return os.WriteFile(targetPath, bytes, 0666)
}

func ReadData(targetPath string) (*library.Catalog, error) {
	bytes, err := os.ReadFile(targetPath)
	if err != nil {
		return nil, err
	}

	var payload FileModel
	if err := json.Unmarshal(bytes, &payload); err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	cat := library.NewCatalog()
	cat.Items = payload.Records
	cat.Members = payload.Users
	cat.SyncBooks(cat.Items) // Автоматически обновит ID

	return cat, nil
}
	