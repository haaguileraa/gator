package config

import(
	"encoding/json"
	"os"
	"path"
)

func getConfigFilePath() (string, error) {
	configDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}
	configFilePath := path.Join(configDir, configFileName)
	return configFilePath, nil
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	body, err := os.ReadFile(configFilePath)

	if err != nil {
		return Config{}, err
	}
	
	var config Config

	if err := json.Unmarshal(body, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
