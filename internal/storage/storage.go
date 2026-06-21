package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/mayankbansal12/yak/internal/model"
)

const filename = "todos.json"
const dataDirEnv = "YAK_DATA_DIR"

type todoStore struct {
	Version int
	Todos   []model.Todo
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

func GetTodos() ([]model.Todo, error) {
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
	return store.Todos, nil
}

func AddTodo(todoItem model.Todo) error {
	todos, err := GetTodos()
	if err != nil {
		return err
	}

	nextItemId := 0
	for _, todo := range todos {
		if todo.ID >= nextItemId {
			nextItemId = todo.ID + 1
		}
	}

	todoItem.ID = nextItemId
	currentTime := time.Now().UTC().Format(time.RFC3339)
	todoItem.CreatedAt = currentTime
	todoItem.UpdatedAt = currentTime

	todos = append(todos, todoItem)
	return SaveTodoToLocalStorage(todos)
}

func SaveTodoToLocalStorage(todos []model.Todo) error {
	store := todoStore{Version: 1, Todos: todos}

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

func MarkTodo(id int) error {
	todos, err := GetTodos()
	if err != nil {
		return err
	}

	for idx, todo := range todos {
		if todo.ID == id {
			currentTime := time.Now().UTC().Format(time.RFC3339)
			todos[idx].CompletedAt = currentTime
			todos[idx].UpdatedAt = currentTime
			return SaveTodoToLocalStorage(todos)
		}
	}

	return fmt.Errorf("Todo with id: %v not found\n", id)
}

func UpdateTodo(id int, updatedItem model.Todo) error {
	todos, err := GetTodos()
	if err != nil {
		return err
	}

	for idx, todo := range todos {
		if todo.ID == id {
			if updatedItem.Title != "" {
				todos[idx].Title = updatedItem.Title
			}
			if updatedItem.Desc != "" {
				todos[idx].Desc = updatedItem.Desc
			}
			if updatedItem.Priority != nil {
				todos[idx].Priority = updatedItem.Priority
			}
			if updatedItem.DueDate != "" {
				todos[idx].DueDate = updatedItem.DueDate
			}
			todos[idx].UpdatedAt = time.Now().UTC().Format(time.RFC3339)
			return SaveTodoToLocalStorage(todos)
		}
	}

	return fmt.Errorf("Todo with id: %v not found\n", id)
}

func DeleteTodo(id int) error {
	todos, err := GetTodos()
	if err != nil {
		return err
	}

	for idx, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:idx], todos[idx+1:]...)
			return SaveTodoToLocalStorage(todos)
		}
	}

	return fmt.Errorf("Todo with id: %v not found\n", id)
}
