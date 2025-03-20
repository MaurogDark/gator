// by convention, we name our package the same as the directory
package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DB_Url    string `json:"db_url"`
	User_Name string `json:"current_user_name"`
}

func (cfg Config) SetUser(user string) Config {
	cfg.User_Name = user
	return cfg
}

func Read() (Config, error) {
	ret := Config{"", ""}
	home, err := os.UserHomeDir()
	if err != nil {
		print("home dir not found!")
		return ret, err
	}
	cfg_path := home + "/" + configFileName
	cfg_file, err := os.Open(cfg_path)
	if err != nil {
		print("could not open config!")
		return ret, err
	}

	decoder := json.NewDecoder(cfg_file)
	if err := decoder.Decode(&ret); err != nil {
		print("could not parse config!")
		return ret, err
	}
	return ret, nil
}

func Write(cfg Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		print("home dir not found!")
		return err
	}
	cfg_path := home + "/" + configFileName
	cfg_file, err := os.Open(cfg_path)
	if err != nil {
		print("could not open config!")
		return err
	}

	encoder := json.NewEncoder(cfg_file)
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		print("could not marshal config!")
		return err
	}

	if err := encoder.Encode(&jsonData); err != nil {
		print("could not encode config!")
		return err
	}
	return nil
}
