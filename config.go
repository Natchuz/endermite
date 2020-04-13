package main

import (
	"encoding/json"
	"github.com/kirsle/configdir"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	configName = "settings.json"
)

var (
	config = EndermiteConfig{
		Accounts: make(map[string]Account),
	}
	configPath = configdir.LocalConfig("endermite")
	configFile = filepath.Join(configPath, configName)
)

type EndermiteConfig struct {
	ClientToken     string             `json:"clientToken"`
	Accounts        map[string]Account `json:"accounts"`
	SelectedAccount string             `json:"selectedAccount"`
}

type Account struct {
	AuthToken string `json:"authToken"`
	Nick      string `json:"nick"`
}

func loadConfig() error {
	if err := configdir.MakePath(configPath); err != nil {
		return err
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		if err := ioutil.WriteFile(configFile, []byte{1}, 0644); err != nil {
			return err
		}
		return updateConfig()
	} else {
		file, err := os.Open(configFile)
		if err != nil {
			return err
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		return decoder.Decode(&config)
	}
}

func updateConfig() error {
	fh, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer fh.Close()
	encoder := json.NewEncoder(fh)
	encoder.SetIndent("", "    ") // pretty print
	return encoder.Encode(&config)
}
