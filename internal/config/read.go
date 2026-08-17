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

	file, err := os.Open(configFilePath)

	if err != nil {
		return Config{}, err
	}
	
	defer file.Close()

	decoder := json.NewDecoder(file)

	var cfg Config

	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
