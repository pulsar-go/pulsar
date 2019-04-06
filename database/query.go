package database

import (
	"strings"

	"github.com/jinzhu/gorm"
)

// All finds all records of a given model
func All(models interface{}) *DB {
	return db.Find(models)
}

// All finds all records of a given model
func (d *DB) All(models interface{}) *DB {
	return All(models)
}

// First finds the first record of the given model
func First(model interface{}, where ...interface{}) *DB {
	return NewDB(db.Lib.First(model, where...))
}

// Where adds a condition to the query statement
func Where(query interface{}, args ...interface{}) *DB {
	if !strings.Contains(query.(string), "= ?") {
		query = query.(string) + " = ?"
	}

	return NewDB(db.Lib.Where(query, args...))
}

// WhereNot adds a condition to the query statement
func WhereNot(query interface{}, args ...interface{}) *DB {
	return NewDB(db.Lib.Not(query, args...))
}

// Take adds a limit to the query
func Take(take interface{}) *gorm.DB {
	return db.Lib.Limit(take)
}

// OrWhere adds an or filter to the query
func (d *DB) OrWhere(query interface{}, args ...interface{}) *DB {
	return NewDB(db.Lib.Or(query, args...))
}

// Find records of a given model
func (d *DB) Find(models interface{}, where ...interface{}) *DB {
	return NewDB(d.Lib.Find(models, where...))
}

// Count 's the database records in the given table
func (d *DB) Count() interface{} {
	var count interface{}
	d.Lib.Count(&count)

	return count
}
