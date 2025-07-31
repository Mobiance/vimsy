package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func EditConfig(name string) error {
	if name == "" {
		var err error
		name, err = CurrentConfig()
		if err != nil {
			return fmt.Errorf("no active config set: %w", err)
		}
	}

	vimsyDir, err := GetVimsyDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(vimsyDir, name)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return errors.New("config does not exist")
	}

	editor := "nvim"
	var targetPath string

	// Try to open init.lua or init.vim if it exists
	initLua := filepath.Join(configPath, "init.lua")
	initVim := filepath.Join(configPath, "init.vim")

	switch {
	case fileExists(initLua):
		targetPath = initLua
	case fileExists(initVim):
		targetPath = initVim
	default:
		// Fallback: open whole directory
		cmd := exec.Command(editor)
		cmd.Dir = configPath
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	cmd := exec.Command(editor, targetPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
