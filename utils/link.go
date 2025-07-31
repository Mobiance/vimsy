package utils

import (
	"fmt"
	"os"
)

func AddConfig(configName string) error {
	target := `C:\Users\shubh\.config\vimsy\` + configName
	link := `C:\Users\shubh\AppData\Local\nvim`

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
