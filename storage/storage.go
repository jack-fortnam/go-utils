package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Address Storage

type Storage struct {
	appName  string
	Config   string
	Data     string
	Cache    string
	Database string
	Logs     string
}

func New(appName string) (*Storage, error) {
	if strings.TrimSpace(appName) == "" || appName == "." || appName == ".." || strings.ContainsAny(appName, `/\`) {
		return nil, fmt.Errorf("invalid application name: %q", appName)
	}

	s := &Storage{appName: appName}

	var err error

	if s.Config, err = s.configDir(); err != nil {
		return nil, err
	}

	if s.Data, err = s.dataDir(); err != nil {
		return nil, err
	}

	if s.Cache, err = s.cacheDir(); err != nil {
		return nil, err
	}

	if s.Logs, err = s.logDir(); err != nil {
		return nil, err
	}

	s.Database = filepath.Join(s.Data, "data.db")

	return s, nil
}

func (s *Storage) configDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find config dir: %w", err)
	}

	return s.createDir(base)
}

func (s *Storage) dataDir() (string, error) {
	var base string

	switch runtime.GOOS {
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("find home dir: %w", err)
		}
		base = filepath.Join(home, "Library", "Application Support")

	case "windows":
		base, _ = os.UserCacheDir()
		if base == "" {
			return "", fmt.Errorf("find local application data dir")
		}

	default:
		base = os.Getenv("XDG_DATA_HOME")
		if base == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", fmt.Errorf("find home dir: %w", err)
			}
			base = filepath.Join(home, ".local", "share")
		}
	}

	return s.createDir(base)
}

func (s *Storage) cacheDir() (string, error) {
	base, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("find cache dir: %w", err)
	}

	return s.createDir(base)
}

func (s *Storage) logDir() (string, error) {
	return s.createDir(filepath.Join(s.Data, "logs"))
}

func (s *Storage) createDir(base string) (string, error) {
	dir := filepath.Join(base, s.appName)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("create directory %s: %w", dir, err)
	}

	return dir, nil
}
