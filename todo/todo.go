// Package todo persists a todo list as JSON.
package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// SaveTodos creates or overwrites path with the JSON list.
func SaveTodos(path string, todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("encode todos: %w", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write todos %q: %w", path, err)
	}
	return nil
}

// LoadTodos reads one JSON list and preserves the underlying error.
func LoadTodos(path string) ([]Todo, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read todos %q: %w", path, err)
	}
	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, fmt.Errorf("decode todos %q: %w", path, err)
	}
	return todos, nil
}
