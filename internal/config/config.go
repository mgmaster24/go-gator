package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Error getting user's home directory")
	}

	file, err := os.Open(homeDir + "/.gatorconfig.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Config{}
	}

	defer file.Close()

	// Decode the JSON data into the struct
	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Config{}
	}

	return config
}

func (config *Config) SetUser(userName string) error {
	config.CurrentUserName = userName
	err := write(*config)
	if err != nil {
		return fmt.Errorf("Error writing to file. Err: %e", err)
	}

	fmt.Println("User config written to $HOME/gatorconfig.json")
	return nil
}

func write(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ") // MarshalIndent for pretty-printing
	if err != nil {
		return err
	}

	filepath, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Error getting user's home directory")
		return err
	}

	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/.gatorconfig.json", nil
}
