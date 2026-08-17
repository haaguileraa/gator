package config

import(
	"encoding/json"
	"os"
)


func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}


func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)	
	return err
} 
