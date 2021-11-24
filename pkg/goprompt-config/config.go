package goprompt_config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
)

const defaultConfigPath = ".config/softteam/goprompt/goprompt.json"

// Config : The main Config type
type Config struct {
	Prefix   string          `json:"main-prefix"`
	Sections []ConfigSection `json:"Sections"`
	Suffix   string          `json:"main-suffix"`
}

type ConfigSection struct {
	SectionType      string `json:"type"`
	FgColor          string `json:"fg-color"`
	BgColor          string `json:"bg-color"`
	Styles           string `json:"styles"`
	Template         string `json:"template"`
	Format           string `json:"format"`
	Separator        string `json:"sep"`
	RemoveIfEmpty    bool   `json:"remove-if-empty"`
	SeparatorFgColor string `json:"sep-fg-color"`
	SeparatorBgColor string `json:"sep-bg-color"`
	SeparatorStyles  string `json:"sep-styles"`
	SeparatorPrefix  string `json:"sep-prefix"`
	SeparatorSuffix  string `json:"sep-suffix"`
}

// NewConfig : Create a new Config
func NewConfig() *Config {
	return &Config{}
}

// Load : Loads the configuration file
func (config *Config) Load() (err error) {
	// Get the path to the Config file
	configPath, err := config.getConfigPath()
	if err != nil {
		fmt.Printf("failed to get the path to the config file\n")
		return err
	}

	// Open Config file
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("failed to open config file: %s\n", configPath)
		return err
	}

	// Parse the JSON document
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		fmt.Printf("failed to parse config file: %s\n", configPath)
		return err
	}

	err = configFile.Close()
	if err != nil {
		fmt.Printf("failed to close config file: %s\n", configPath)
		return err
	}

	return nil
}

//
// Private functions
//

// getConfigPath gets the path to the Config file
func (config *Config) getConfigPath() (string, error) {
	home, err := config.getHomeDirectory()
	if err != nil {
		return "", err
	}

	return path.Join(home, defaultConfigPath), nil
}

// getHomeDirectory gets the current users home directory
func (config *Config) getHomeDirectory() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory : %s", err)
	}
	return u.HomeDir, nil
}
