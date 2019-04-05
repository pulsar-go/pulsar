package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pulsar-go/pulsar/database"
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
	Database database.DBconfigs `toml:"database"`
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
