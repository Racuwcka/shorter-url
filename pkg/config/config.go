package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type EnvFile struct {
	data map[string]string
}

func LoadEnvFile(filename string) (*EnvFile, error) {
	env := &EnvFile{
		data: make(map[string]string),
	}

	file, err := os.Open(filename)
	if err != nil {
		return env, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.ToLower(strings.TrimSpace(parts[0]))
			value := strings.TrimSpace(parts[1])
			env.data[key] = value
		}
	}

	return env, nil
}

func (e *EnvFile) GetEnvString(key string, defaultValue string) string {
	if value, exists := e.data[key]; exists {
		return value
	}
	return defaultValue
}

func (e *EnvFile) GetEnvInt(key string, defaultValue int) int {
	if value, exists := e.data[key]; exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
