package db

import (
	"strings"
)

// Builder represents the current database used.
var Builder *DB

// Create the model on the database
func (b *DB) Create(model interface{}) *DB {
	return SetLib(b.lib.Create(model))
}

// Save the model to the database
func (b *DB) Save(model interface{}) *DB {
	return SetLib(b.lib.Save(model))
}

// Delete an existing model
func (b *DB) Delete(model interface{}, where ...interface{}) *DB {
	return SetLib(b.lib.Delete(model, where))
}

// All finds all records of a given model
func (b *DB) All(models interface{}) *DB {
	return b.Find(models)
}

// First finds the first record of the given model
func (b *DB) First(model interface{}, where ...interface{}) *DB {
	return SetLib(b.lib.First(model, where...)).clone()
}

// Where adds a condition to the query statement
func (b *DB) Where(query interface{}, args ...interface{}) *DB {
	if !strings.Contains(query.(string), "= ?") {
		query = query.(string) + " = ?"
	}

	return SetLib(b.lib.Where(query, args...)).clone()
}

// WhereNot adds a condition to the query statement
func (b *DB) WhereNot(query interface{}, args ...interface{}) *DB {
	return SetLib(b.lib.Not(query, args...)).clone()
}

// Take adds a limit to the query
func (b *DB) Take(take interface{}) *DB {
	return SetLib(b.lib.Limit(take)).clone()
}

// OrWhere adds an or filter to the query
func (b *DB) OrWhere(query interface{}, args ...interface{}) *DB {
	return SetLib(b.lib.Or(query, args...)).clone()
}

// Find records of a given model
func (b *DB) Find(models interface{}, where ...interface{}) *DB {
	return SetLib(b.lib.Find(models, where...)).clone()
}

// Count 's the database records in the given table
func (b *DB) Count() interface{} {
	var count interface{}
	b.lib.Count(&count)

	return count
}
