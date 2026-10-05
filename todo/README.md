# Task 1 — Todo-list persistence (JSON)

Implement `SaveTodos` and `LoadTodos` in `todo.go`.

**Requirements**

- `SaveTodos(path string, todos []Todo) error` writes the list to `path`
  as JSON, creating or overwriting the file.
- `LoadTodos(path string) ([]Todo, error)` reads and parses the file at
  `path`, returning a clear error for a missing file or malformed JSON.
- At least **80% statement coverage** for this package
  (`go test -cover ./todo/...`).

**Do not edit `todo_test.go`** — it's the specification for this task.
Run it locally with:

```bash
go test -v ./todo/...
go test -cover ./todo/...
```
