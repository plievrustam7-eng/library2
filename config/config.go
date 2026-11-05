package config

import "errors"

func GetPort(cfg map[string]string) (string, error) {
	port, ok := cfg["PORT"]
	if !ok {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	return port, nil
}