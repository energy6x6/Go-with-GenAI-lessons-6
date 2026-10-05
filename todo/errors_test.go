package todo

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestPersistenceErrors(t *testing.T) {
	t.Run("marshal error", func(t *testing.T) {
		err := SaveTodos(filepath.Join(t.TempDir(), "todos.json"), []Todo{{CreatedAt: time.Date(10000, 1, 1, 0, 0, 0, 0, time.UTC)}})
		var marshalErr *json.MarshalerError
		if !errors.As(err, &marshalErr) {
			t.Errorf("error = %v, want wrapped marshaler error", err)
		}
	})
	t.Run("write error", func(t *testing.T) {
		err := SaveTodos(filepath.Join(t.TempDir(), "missing", "todos.json"), nil)
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("error = %v, want wrapped not-exist", err)
		}
	})
	t.Run("missing file preserves cause", func(t *testing.T) {
		_, err := LoadTodos(filepath.Join(t.TempDir(), "missing.json"))
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("error = %v, want wrapped not-exist", err)
		}
	})
	t.Run("JSON error preserves cause", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "bad.json")
		if err := os.WriteFile(path, []byte("bad"), 0o600); err != nil {
			t.Fatal(err)
		}
		_, err := LoadTodos(path)
		var syntaxErr *json.SyntaxError
		if !errors.As(err, &syntaxErr) {
			t.Errorf("error = %v, want wrapped syntax error", err)
		}
	})
}
