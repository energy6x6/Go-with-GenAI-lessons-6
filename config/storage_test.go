package config

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

// mockStorage is handwritten and uses no mocking framework.
type mockStorage struct {
	file               *mockFile
	openErr, createErr error
}

func (s mockStorage) Open(string) (io.ReadCloser, error)    { return s.file, s.openErr }
func (s mockStorage) Create(string) (io.WriteCloser, error) { return s.file, s.createErr }

type mockFile struct {
	bytes.Buffer
	readErr, writeErr, closeErr error
	closed                      bool
}

func (f *mockFile) Read(p []byte) (int, error) {
	if f.readErr != nil {
		return 0, f.readErr
	}
	return f.Buffer.Read(p)
}
func (f *mockFile) Write(p []byte) (int, error) {
	if f.writeErr != nil {
		return 0, f.writeErr
	}
	return f.Buffer.Write(p)
}
func (f *mockFile) Close() error { f.closed = true; return f.closeErr }

func TestStorageFailures(t *testing.T) {
	diskFull := errors.New("disk full")
	closeFailure := errors.New("close failed")
	cases := []struct {
		name                                            string
		save                                            bool
		openErr, createErr, readErr, writeErr, closeErr error
		want                                            error
	}{
		{name: "permission denied opening", openErr: os.ErrPermission, want: os.ErrPermission},
		{name: "permission denied creating", save: true, createErr: os.ErrPermission, want: os.ErrPermission},
		{name: "read failure", readErr: io.ErrUnexpectedEOF, want: io.ErrUnexpectedEOF},
		{name: "disk full writing", save: true, writeErr: diskFull, want: diskFull},
		{name: "close after read", closeErr: closeFailure, want: closeFailure},
		{name: "close after write", save: true, closeErr: closeFailure, want: closeFailure},
		{name: "write and close fail", save: true, writeErr: diskFull, closeErr: closeFailure, want: diskFull},
		{name: "read and close fail", readErr: io.ErrUnexpectedEOF, closeErr: closeFailure, want: io.ErrUnexpectedEOF},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := &mockFile{readErr: tc.readErr, writeErr: tc.writeErr, closeErr: tc.closeErr}
			f.Buffer.WriteString(`{"server_port":8080,"environment":"development"}`)
			store := NewStore(mockStorage{file: f, openErr: tc.openErr, createErr: tc.createErr})
			var err error
			if tc.save {
				err = store.SaveConfig("config.json", Config{})
			} else {
				_, err = store.LoadConfig("config.json")
			}
			if !errors.Is(err, tc.want) {
				t.Fatalf("error = %v, want wrapped %v", err, tc.want)
			}
			if !strings.Contains(err.Error(), "config.json") {
				t.Errorf("path context missing: %v", err)
			}
			if tc.closeErr != nil && !errors.Is(err, tc.closeErr) {
				t.Errorf("close error lost: %v", err)
			}
			wantClosed := tc.openErr == nil && tc.createErr == nil
			if f.closed != wantClosed {
				t.Errorf("closed = %v, want %v", f.closed, wantClosed)
			}
		})
	}
}
