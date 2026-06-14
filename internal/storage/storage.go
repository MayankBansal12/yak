package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"yak-cli/internal/model"
)

const filename = "todos.json"
const dataDirEnv = "YAK_DATA_DIR"

type todoStore struct {
	version int
	todos []model.Todo
}

func GetStorageFilePath() (string, error) {
	var userDir string

	if v := os.Getenv(dataDirEnv); v != "" {
		userDir = v
	} else if v := os.Getenv("XDG_DATA_HOME"); v != "" {
		userDir = v
	} else {
		homeDir, err := os.UserHomeDir()
        if err != nil {
            return "", err
        }
        userDir = filepath.Join(homeDir, ".local", "share", "yak")
	}

	// ensure dir exists before joining filename
	if err := os.MkdirAll(userDir, 0o755); err != nil {
		return "", err
	}

	return filepath.Join(userDir, filename), nil
}

func LoadTodo() ([]model.Todo, error) {
	todoFilePath, err := GetStorageFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(todoFilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []model.Todo{}, nil
		}
		return nil, err
	}

	var store todoStore
	err = json.Unmarshal(data, &store)
	if err != nil {
		backupPath := todoFilePath + ".bak"
		if err := os.Rename(todoFilePath, backupPath); err != nil {
		    fmt.Fprintln(os.Stderr, "Warning: unable to create backup for corrupt todos.json. Starting with empty list.")
		} else {
		    fmt.Fprintf(os.Stderr, "Warning: corrupt todos.json moved to %s. Starting with empty list.\n", backupPath)
		}
		return []model.Todo{}, nil
	}
	return store.todos, nil
}

func SaveTodo(todos []model.Todo) error {
	store := todoStore{version: 1, todos: todos}

	storeJson, err := json.Marshal(store)
	if err != nil {
		return err
	}

	todoFilePath, err := GetStorageFilePath()
	if err != nil {
		return err
	}

	tmpFilePath := todoFilePath + ".tmp"
	err = os.WriteFile(tmpFilePath, storeJson, 0o644)
	if err != nil {
		return err
	}

	err = os.Rename(tmpFilePath, todoFilePath)
	if err != nil {
		_ = os.Remove(tmpFilePath)
		return err
	}

	return nil
}
