package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	// The following imports are for the database drivers.
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DBconfigs represents a global database settings.
type DBconfigs struct {
	Driver      string `toml:"driver"`
	Database    string `toml:"database"`
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	AutoMigrate bool   `toml:"auto_migrate"`
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
func Open(config *DBconfigs) {
	// Create the arguments
	var args string
	switch config.Driver {
	case "mysql":
		args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Database)
	case "postgres":
		args = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.Host, config.Port, config.User, config.Database, config.Password)
	case "sqlite3":
		args = config.Database
	default:
		log.Fatalf("Database driver '%s' is not supported.\n", config.Driver)
	}
	// Open the database
	dbOpened, err := gorm.Open(config.Driver, args)
	if err != nil {
		log.Fatalln(err)
	}
	DB = dbOpened
}
