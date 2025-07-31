package utils

import (
	"os"
	"path/filepath"
)

func DeleteConfig(configName string) error {
	vimsyDir, err := GetVimsyDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(vimsyDir, configName)

	// Check if the config directory exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return os.ErrNotExist // Config does not exist
	} else if err != nil {
		return err // Other error occurred
	}

	// Remove the config directory
	if err := os.RemoveAll(configPath); err != nil {
		return err // Failed to remove the config directory
	}

	return nil // Successfully deleted the config
}
