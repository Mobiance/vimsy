package utils

import (
	"os"
	"path/filepath"
)

func RenameConfig(oldName, newName string) error {
	// Get the Vimsy directory
	vimsyDir, err := GetVimsyDir()
	if err != nil {
		return err
	}

	// Define the old and new config paths
	oldConfigPath := filepath.Join(vimsyDir, oldName)
	newConfigPath := filepath.Join(vimsyDir, newName)

	// Check if the old config exists
	if _, err := os.Stat(oldConfigPath); os.IsNotExist(err) {
		return os.ErrNotExist // Old config does not exist
	} else if err != nil {
		return err // Other error occurred
	}

	// Rename the config directory
	if err := os.Rename(oldConfigPath, newConfigPath); err != nil {
		return err // Failed to rename the config directory
	}

	return nil // Successfully renamed the config
}
