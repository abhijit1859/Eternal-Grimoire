package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Localstorage struct {
	dir string
}

func (ls *Localstorage) NewLocalStorage(dir string) *Localstorage {
	return &Localstorage{
		dir: dir,
	}
}

func (ls *Localstorage) Store(ctx context.Context, name string, data io.Reader) error {
	if name == "" || filepath.Base(name) != name {
		return fmt.Errorf("Invalid backup name")
	}

	if err := os.MkdirAll(ls.dir, 0755); err != nil {
		return fmt.Errorf("create backup directory: %w", err)
	}

	path := filepath.Join(ls.dir, name)

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create backup file: %w", err)
	}

	_, copyErr := io.Copy(file, data)
	closeErr := file.Close()

	if copyErr != nil {
		return fmt.Errorf("write backup file: %w", copyErr)
	}
	if closeErr != nil {
		return fmt.Errorf("close backup file: %w", closeErr)
	}

	return nil
}

func (ls *Localstorage) Retrieve(ctx context.Context, name string) (io.ReadCloser, error) {
	if name == "" || filepath.Base(name) != name {
		return nil, fmt.Errorf("Invalid backup name")
	}

	file, err := os.Open(filepath.Join(ls.dir, name))

	if err != nil {
		return nil, fmt.Errorf("open backup file: %w", err)
	}

	return file, nil
}

func (ls *Localstorage) List(ctx context.Context) ([]string, error) {
	enteries, err := os.ReadDir(ls.dir)

	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("read backup directory: %w", err)
	}

	var names []string

	for _, entry := range enteries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	return names, nil
}
