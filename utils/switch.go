package utils

import (
	"fmt"
	"os"
)

func SwitchConfig(configName string) error {
	target := os.Getenv("VIMSY_CONFIG_DIR") + `\` + configName
	link := os.Getenv("NEOVIM_CONFIG_DIR")

	if !IsAdmin() {
		return fmt.Errorf("⚠️  Not running as admin — symlink creation may fail.")
	}
	err := RemoveIfExists(link)
	if err != nil {
		return fmt.Errorf("Failed to remove existing path: %v", err)
	}
	err = os.Symlink(target, link)
	if err != nil {
		return fmt.Errorf("Symlink failed: %v", err)
	} else {
		return nil
	}
}
