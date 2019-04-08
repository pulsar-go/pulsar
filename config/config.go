package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Config represents the pulsar server settings structure.
type Config struct {
	Server struct {
		Host        string `toml:"host"`
		Port        string `toml:"port"`
		Development bool   `toml:"development"`
	} `toml:"server"`
	HTTPS struct {
		Enabled  bool   `toml:"enabled"`
		CertFile string `toml:"cert_file"`
		KeyFile  string `toml:"key_file"`
	} `toml:"https"`
	Views struct {
		Path string `toml:"path"`
	} `toml:"views"`
	Database struct {
		Driver      string `toml:"driver"`
		Database    string `toml:"database"`
		Host        string `toml:"host"`
		Port        string `toml:"port"`
		User        string `toml:"user"`
		Password    string `toml:"password"`
		AutoMigrate bool   `toml:"auto_migrate"`
	} `toml:"database"`
	Mail struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Identity string `toml:"identity"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		From     string `toml:"from"`
	} `toml:"mail"`
	Queue struct {
		Routines string `toml:"routines"`
	} `toml:"queue"`
}

// Settings define the global settings for pulsar.
var Settings Config

// Set sets the configuration from a configuration file.
func Set(path string) {
	// Open the server configuration file.
	absPath, _ := filepath.Abs(filepath.Clean(path))
	if _, err := toml.DecodeFile(absPath, &Settings); err != nil {
		log.Fatalln("There was an error decoding file " + absPath)
	}
	// Transform the relative paths into absolute.
	Settings.HTTPS.CertFile, _ = filepath.Abs(filepath.Dir(path) + "/" + filepath.Clean(Settings.HTTPS.CertFile))
	Settings.HTTPS.KeyFile, _ = filepath.Abs(filepath.Dir(path) + "/" + filepath.Clean(Settings.HTTPS.KeyFile))
}
