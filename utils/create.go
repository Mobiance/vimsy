package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func CreateConfig(name string) error {
	vimsyDir, err := GetVimsyDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(vimsyDir, name)

	if _, err := os.Stat(configPath); err == nil {
		return errors.New("config already exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(configPath, 0755); err != nil {
		return err
	}

	// Optional: Create a starter init.lua or README.md
	initPath := filepath.Join(configPath, "init.lua")
	if err := os.WriteFile(initPath, []byte("-- Your Neovim config starts here\n"), 0644); err != nil {
		return err
	}

	return nil
}
