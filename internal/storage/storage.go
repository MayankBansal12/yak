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

func LocalTodoFilePath() (string, error) {
	userDir := os.Getenv(dataDirEnv)
	if userDir == "" {
		baseDir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}
		userDir = filepath.Join(baseDir, "yak", filename)
	}
	err := os.MkdirAll(userDir, 0o755)
	if(err != nil){
		return "", err
	}
	return userDir, nil
}

func LoadTodo() ([]model.Todo, error) {
	todoFilePath, err := LocalTodoFilePath()
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
		    fmt.Fprintln(os.Stderr, "Warning: unable to create backup for corrupt todos.json...Starting with empty list.")
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

	todoFilePath, err := LocalTodoFilePath()
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
