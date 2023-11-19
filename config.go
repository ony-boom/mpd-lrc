package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	Port            int    `toml:"port"`
	Host            string `toml:"host"`
	MusicPath       string `toml:"musicPath"`
	TitleColor      string `toml:"titleColor"`
	PollingDelay    int    `toml:"pollingDelay"`
	ActiveLineColor string `toml:"activeLineColor"`
	ServerPort      int    `toml:"serverPort"`
}

func getConfig() Config {
	var conf Config

	_, err := toml.DecodeFile(getConfigFilePath(), &conf)

	if err != nil {
		return Config{
			ActiveLineColor: "3",
			TitleColor:      "2",
			PollingDelay:    150,
			ServerPort:      6900,
			Port:            6600,
			Host:            "127.0.0.1",
			MusicPath:       filepath.Join("~/music"),
		}
	}

	return conf
}

func getConfigFilePath() string {
	const filename = "mpdLrc.toml"

	baseConfigDir := strings.TrimSpace(os.Getenv("XDG_CONFIG_HOME"))

	if baseConfigDir == "" {
		systemHome, err := homedir.Dir()
		if err != nil {
			systemHome = "~"
		}
		baseConfigDir = filepath.Join(systemHome, ".config")
	}

	return filepath.Join(baseConfigDir, filename)
}
