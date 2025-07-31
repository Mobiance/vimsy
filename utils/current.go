package utils

import (
	"os"
	"path/filepath"
)

func CurrentConfig() (string, error) {
	// Get the current configuration name from the environment variable
	currentConfig := os.Getenv("NEOVIM_CONFIG_DIR")
	activeConfig, err := os.Readlink(currentConfig)
	if err != nil {
		if os.IsNotExist(err) {
			return "", os.ErrNotExist
		}
		return "", err
	}
	if activeConfig == "" {
		return "", os.ErrNotExist
	}

	// Extract the configuration name from the path
	configName := filepath.Base(activeConfig)

	return configName, nil
}
