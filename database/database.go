package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pulsar-go/pulsar/config"

	// The following imports are for the database drivers.
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Model represents the base database model.
type Model gorm.Model

// DB represents the database structure used
type DB struct {
	Lib *gorm.DB
}

// DB represents the current database used.
var db *DB

// Models stores the current set of application models.
var Models []interface{}

// AddModels add the given models to the model list.
func AddModels(models ...interface{}) {
	Models = append(Models, models...)
}

// Open opens a new database connection.
func Open() {
	// Create the arguments
	var args string
	// Copy to reduce code size.
	s := &config.Settings.Database
	switch s.Driver {
	case "mysql":
		args = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Host, s.Port, s.Database)
	case "postgres":
		args = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", s.Host, s.Port, s.User, s.Database, s.Password)
	case "sqlite3":
		args = s.Database
	default:
		log.Fatalf("Database driver '%s' is not supported.\n", s.Driver)
	}
	// Open the database
	dbOpened, err := gorm.Open(s.Driver, args)
	if err != nil {
		log.Fatalln(err)
	}
	db = NewDB(dbOpened)
}

// Lib returns the underlying library instance
func Lib() *gorm.DB {
	return db.Lib
}

// NewDB converts a Lib response into a DB response
func NewDB(db *gorm.DB) *DB {
	newDB := new(DB)
	newDB.Lib = db
	return newDB
}
