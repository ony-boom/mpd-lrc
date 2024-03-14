package main

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
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
	musicPath, err := filepath.Abs(conf.MusicPath)

	if err != nil {
		musicPath = "~/Music"
	}

	if err != nil {
		return Config{
			ActiveLineColor: "3",
			TitleColor:      "2",
			PollingDelay:    150,
			ServerPort:      6900,
			Port:            6600,
			Host:            "127.0.0.1",
			MusicPath:      musicPath, 
		}
	}

	return conf
}

func getConfigFilePath() string {
	const filename = "mpdLrc.toml"

	baseConfigDir := strings.TrimSpace(os.Getenv("XDG_CONFIG_HOME"))

	if baseConfigDir == "" {
		systemHome, _ := user.Current()
		baseConfigDir = filepath.Join(systemHome.HomeDir, ".config")
	}

	return filepath.Join(baseConfigDir, filename)
}
