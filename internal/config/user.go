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

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	
	err = os.WriteFile(configFilePath, data, 0777)
	return err
} 
