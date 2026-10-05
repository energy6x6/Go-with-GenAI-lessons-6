// Package config manages JSON server configuration files.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Config struct {
	ServerPort    int    `json:"server_port"`
	Environment   string `json:"environment"`
	DatabaseURL   string `json:"database_url,omitempty"`
	DebugMode     bool   `json:"debug_mode"`
	AdminPassword string `json:"-"`
}

// FileStorage abstracts opening files for reading and creating them for writing.
// Returned streams belong to the caller, which must close them.
type FileStorage interface {
	Open(path string) (io.ReadCloser, error)
	Create(path string) (io.WriteCloser, error)
}

type diskStorage struct{}

func (diskStorage) Open(path string) (io.ReadCloser, error)    { return os.Open(path) }
func (diskStorage) Create(path string) (io.WriteCloser, error) { return os.Create(path) }

// Store allows a storage implementation to be supplied without global state.
type Store struct{ storage FileStorage }

func NewStore(storage FileStorage) *Store { return &Store{storage: storage} }

// SaveConfig writes an indented JSON file using the local filesystem.
func SaveConfig(path string, cfg Config) error {
	return NewStore(diskStorage{}).SaveConfig(path, cfg)
}

// LoadConfig reads a JSON file using the local filesystem.
func LoadConfig(path string) (Config, error) {
	return NewStore(diskStorage{}).LoadConfig(path)
}

// SaveConfig streams JSON and preserves creation, write and close errors.
func (s *Store) SaveConfig(path string, cfg Config) (err error) {
	f, err := s.storage.Create(path)
	if err != nil {
		return fmt.Errorf("create config %q: %w", path, err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close config %q: %w", path, closeErr))
		}
	}()
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("write config %q: %w", path, err)
	}
	return nil
}

// LoadConfig accepts exactly one JSON value, matching Unmarshal's behavior.
func (s *Store) LoadConfig(path string) (cfg Config, err error) {
	f, err := s.storage.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("open config %q: %w", path, err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = errors.Join(err, fmt.Errorf("close config %q: %w", path, closeErr))
		}
	}()
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("decode config %q: %w", path, err)
	}
	var extra json.RawMessage
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return Config{}, fmt.Errorf("decode config %q: multiple JSON values", path)
		}
		return Config{}, fmt.Errorf("decode config %q trailing data: %w", path, err)
	}
	return cfg, nil
}

// ValidateConfig checks the port and exact, case-sensitive environment name.
func ValidateConfig(cfg Config) error {
	if cfg.ServerPort < 1024 || cfg.ServerPort > 65535 {
		return fmt.Errorf("server_port must be between 1024 and 65535: %d", cfg.ServerPort)
	}
	switch cfg.Environment {
	case "development", "staging", "production":
		return nil
	default:
		return fmt.Errorf("invalid environment: %q", cfg.Environment)
	}
}
