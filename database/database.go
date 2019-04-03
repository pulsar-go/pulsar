package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// The following imports are for the database drivers.
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBsettings represents a global database settings.
type DBsettings struct {
	Driver      string `toml:"driver"`
	Database    string `toml:"database"`
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	AutoMigrate bool   `toml:"auto_migrate`
}

// Model represents the base database model.
type Model gorm.Model

// DB represents the current database used.
var DB *gorm.DB

// Models stores the current set of application models.
var Models []interface{}

// AddModels add the given models to the model list.
func AddModels(models ...interface{}) {
	Models = append(Models, models...)
}

// Open opens a new database connection.
func Open(s *DBsettings) {
	// Create the arguments
	var args string
	switch s.Driver {
	case "mysql":
		args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.user, s.Password, s.Host, s.Port, s.Database)
	case "postgres":
		args = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", s.Host, s.Port, s.User, s.Database, s.Password)
	case "sqlite3":
		args = s.Database
	case "mssql":
		args = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", s.Username, s.Password, s.Host, s.Port, s.Database)
	default:
		log.Fatalf("Database driver '%s' is not supported.\n", s.Driver)
	}
	// Open the database
	DB, err := gorm.Open(s.Driver, args)
}
