package utils

import (
	"os"
)

func ListConfigs() ([]string, error) {
	vimsyConfigDir := os.Getenv("VIMSY_CONFIG_DIR")

	if vimsyConfigDir == "" {
		return nil, os.ErrNotExist
	}

	folder, err := os.ReadDir(vimsyConfigDir)
	if err != nil {
		return nil, err
	}

	var configs []string
	for _, file := range folder {
		if file.IsDir() {
			configs = append(configs, file.Name())
		}
	}
	if len(configs) == 0 {
		return nil, os.ErrNotExist
	}

	return configs, nil

}
