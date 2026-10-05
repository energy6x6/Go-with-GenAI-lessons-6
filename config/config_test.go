package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cases := []struct {
		name, data       string
		missing, wantErr bool
		want             Config
	}{
		{name: "successful read", data: `{"server_port":8080,"environment":"development","database_url":"db","debug_mode":true,"admin_password":"ignored"}`, want: Config{ServerPort: 8080, Environment: "development", DatabaseURL: "db", DebugMode: true}},
		{name: "file not found", missing: true, wantErr: true},
		{name: "corrupted JSON", data: `{"server_port":`, wantErr: true},
		{name: "empty file", wantErr: true},
		{name: "wrong field type", data: `{"server_port":"8080"}`, wantErr: true},
		{name: "extra JSON value", data: `{} {}`, wantErr: true},
		{name: "trailing junk", data: `{} bad`, wantErr: true},
		{name: "trailing whitespace", data: "{}\n "},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "config.json")
			if !tc.missing {
				if err := os.WriteFile(path, []byte(tc.data), 0o600); err != nil {
					t.Fatalf("setup: %v", err)
				}
			}
			got, err := LoadConfig(path)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.missing && !errors.Is(err, os.ErrNotExist) {
				t.Errorf("missing error not preserved: %v", err)
			}
			if !tc.wantErr && got != tc.want {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestSaveConfig(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	for _, db := range []string{"postgres://localhost/app", ""} {
		cfg := Config{ServerPort: 8080, Environment: "staging", DatabaseURL: db, DebugMode: true, AdminPassword: "secret"}
		if err := SaveConfig(path, cfg); err != nil {
			t.Fatalf("save: %v", err)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read: %v", err)
		}
		if !strings.Contains(string(data), "\n  \"server_port\"") {
			t.Errorf("JSON is not indented: %s", data)
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(data, &fields); err != nil {
			t.Fatalf("parse: %v", err)
		}
		for _, key := range []string{"server_port", "environment", "debug_mode"} {
			if _, ok := fields[key]; !ok {
				t.Errorf("missing JSON key %s", key)
			}
		}
		if _, ok := fields["database_url"]; ok != (db != "") {
			t.Error("database_url omitempty not respected")
		}
		if strings.Contains(string(data), "secret") || strings.Contains(string(data), "password") {
			t.Error("password serialized")
		}
		got, err := LoadConfig(path)
		if err != nil {
			t.Fatalf("load: %v", err)
		}
		cfg.AdminPassword = ""
		if got != cfg {
			t.Errorf("got %+v, want %+v", got, cfg)
		}
	}
}

func TestSaveConfigMissingDirectory(t *testing.T) {
	err := SaveConfig(filepath.Join(t.TempDir(), "missing", "config.json"), Config{})
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("error = %v, want wrapped not-exist", err)
	}
}
