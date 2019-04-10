package config

import (
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// ServerConfig specifies the configuration for the server file.
type ServerConfig struct {
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	Development bool   `toml:"development"`
}

// CertificateConfig specifies the configuration for the certificate file.
type CertificateConfig struct {
	Enabled  bool   `toml:"enabled"`
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

// ViewsConfig specifies the configuration for the view file.
type ViewsConfig struct {
	Path string `toml:"path"`
}

// DatabaseConfig specifies the configuration for the database file.
type DatabaseConfig struct {
	Driver      string `toml:"driver"`
	Database    string `toml:"database"`
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	AutoMigrate bool   `toml:"auto_migrate"`
}

// MailConfig specifies the configuration for the mail file.
type MailConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Identity string `toml:"identity"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	From     string `toml:"from"`
}

// QueueConfig specifies the configuration for the queue file.
type QueueConfig struct {
	Routines string `toml:"routines"`
}

// Config represents the pulsar server settings structure.
type Config struct {
	Server      ServerConfig
	Certificate CertificateConfig
	Views       ViewsConfig
	Database    DatabaseConfig
	Mail        MailConfig
	Queue       QueueConfig
}

// Settings define the global settings for pulsar.
var Settings Config

// @todo revisit with map[string]interface{} to make it dynamic
func setConfigOf(file string, v interface{}) {
	absPath, _ := filepath.Abs(filepath.Clean("./config/" + file + ".toml"))
	if _, err := toml.DecodeFile(absPath, v); err != nil {
		log.Fatalln("There was an error decoding file " + absPath + ", Error: " + err.Error())
	}
}

// Set sets the configuration from a configuration file.
func init() {
	// Server config
	setConfigOf("server", &Settings.Server)
	// Certificate config
	setConfigOf("certificate", &Settings.Certificate)
	// Views config
	setConfigOf("views", &Settings.Views)
	// Database config
	setConfigOf("database", &Settings.Database)
	// Mail config
	setConfigOf("mail", &Settings.Mail)
	// Queue config
	setConfigOf("queue", &Settings.Queue)
	// Transform the relative paths into absolute.
	Settings.Certificate.CertFile, _ = filepath.Abs(filepath.Dir("./config") + "/" + filepath.Clean(Settings.Certificate.CertFile))
	Settings.Certificate.KeyFile, _ = filepath.Abs(filepath.Dir("./config") + "/" + filepath.Clean(Settings.Certificate.KeyFile))
}
