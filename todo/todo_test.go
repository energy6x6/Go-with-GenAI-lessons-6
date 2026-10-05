// This file is the specification for Task 1. You should not need to edit
// it — implement SaveTodos and LoadTodos in todo.go until every test here
// passes and `go test -cover ./todo/...` reports at least 80%.
//
// Note for students: every case below uses t.Errorf (not t.Fatalf) where
// possible and is wrapped in its own t.Run subtest. That is intentional —
// it means a single missing feature will NOT hide the other failures.
// Run `go test -v ./todo/...` and read the whole log: every FAIL line is
// a separate piece of missing or incorrect behavior.
package todo

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSaveAndLoadTodos_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "todos.json")

	want := []Todo{
		{ID: 1, Title: "Write tests", Done: false, CreatedAt: time.Date(2026, 1, 10, 9, 0, 0, 0, time.UTC)},
		{ID: 2, Title: "Ship feature", Done: true, CreatedAt: time.Date(2026, 1, 11, 9, 0, 0, 0, time.UTC)},
	}

	if err := SaveTodos(path, want); err != nil {
		t.Fatalf("SaveTodos() error = %v, want nil", err)
	}

	got, err := LoadTodos(path)
	if err != nil {
		t.Fatalf("LoadTodos() error = %v, want nil", err)
	}

	if len(got) != len(want) {
		t.Fatalf("LoadTodos() returned %d todos, want %d", len(got), len(want))
	}

	for i := range want {
		i := i
		t.Run(want[i].Title, func(t *testing.T) {
			if got[i].ID != want[i].ID {
				t.Errorf("ID = %d, want %d", got[i].ID, want[i].ID)
			}
			if got[i].Title != want[i].Title {
				t.Errorf("Title = %q, want %q", got[i].Title, want[i].Title)
			}
			if got[i].Done != want[i].Done {
				t.Errorf("Done = %v, want %v", got[i].Done, want[i].Done)
			}
		})
	}
}

func TestSaveAndLoadTodos_TableDriven(t *testing.T) {
	cases := []struct {
		name  string
		todos []Todo
	}{
		{"single item", []Todo{{ID: 1, Title: "one"}}},
		{"multiple items", []Todo{{ID: 1, Title: "one"}, {ID: 2, Title: "two", Done: true}}},
		{"empty list", []Todo{}},
		{"nil list", nil},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "todos.json")

			if err := SaveTodos(path, tc.todos); err != nil {
				t.Fatalf("SaveTodos() error = %v, want nil", err)
			}

			got, err := LoadTodos(path)
			if err != nil {
				t.Fatalf("LoadTodos() error = %v, want nil", err)
			}
			if len(got) != len(tc.todos) {
				t.Errorf("round trip returned %d todos, want %d", len(got), len(tc.todos))
			}
		})
	}
}

func TestSaveTodos_CreatesFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "created.json")

	if _, err := os.Stat(path); err == nil {
		t.Fatalf("test setup error: %s already exists", path)
	}

	if err := SaveTodos(path, []Todo{{ID: 1, Title: "x"}}); err != nil {
		t.Fatalf("SaveTodos() error = %v, want nil", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Errorf("expected file %s to exist after SaveTodos, stat error = %v", path, err)
	}
}

func TestSaveTodos_OverwritesExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "todos.json")

	if err := SaveTodos(path, []Todo{{ID: 1, Title: "first version"}}); err != nil {
		t.Fatalf("SaveTodos() first call error = %v, want nil", err)
	}
	if err := SaveTodos(path, []Todo{{ID: 2, Title: "second version"}}); err != nil {
		t.Fatalf("SaveTodos() second call error = %v, want nil", err)
	}

	got, err := LoadTodos(path)
	if err != nil {
		t.Fatalf("LoadTodos() error = %v, want nil", err)
	}
	if len(got) != 1 {
		t.Fatalf("got %d todos after overwrite, want 1", len(got))
	}
	if got[0].Title != "second version" {
		t.Errorf("Title = %q, want %q (file was not overwritten)", got[0].Title, "second version")
	}
}

func TestLoadTodos_MissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "does-not-exist.json")

	_, err := LoadTodos(path)
	if err == nil {
		t.Fatal("LoadTodos() error = nil, want an error for a missing file")
	}
}

func TestLoadTodos_MalformedJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")

	if err := os.WriteFile(path, []byte("{not valid json"), 0o644); err != nil {
		t.Fatalf("test setup error: %v", err)
	}

	_, err := LoadTodos(path)
	if err == nil {
		t.Fatal("LoadTodos() error = nil, want an error for malformed JSON")
	}
}

func TestLoadTodos_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.json")

	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		t.Fatalf("test setup error: %v", err)
	}

	// An empty file is malformed JSON (not even "[]"), so this should
	// also return an error rather than an empty, successful result.
	_, err := LoadTodos(path)
	if err == nil {
		t.Fatal("LoadTodos() error = nil, want an error for an empty file")
	}
}
