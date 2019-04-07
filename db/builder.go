package db

import (
	"strings"
)

// All finds all records of a given model
func (b *DB) All(models interface{}) *DB {
	return b.clone(b.DB.Find(models))
}

// Where adds a condition to the query statement
func (b *DB) Where(query interface{}, args ...interface{}) *DB {
	if !strings.Contains(query.(string), "= ?") {
		query = query.(string) + " = ?"
	}

	return b.clone(b.DB.Where(query, args...))
}

// WhereNot adds a condition to the query statement
func (b *DB) WhereNot(query interface{}, args ...interface{}) *DB {
	return b.clone(b.Not(query, args...))
}

// Take adds a limit to the query
func (b *DB) Take(take interface{}) *DB {
	return b.clone(b.Limit(take))
}

// OrWhere adds an or filter to the query
func (b *DB) OrWhere(query interface{}, args ...interface{}) *DB {
	return b.clone(b.Or(query, args...))
}

// Count 's the database records in the given table
func (b *DB) Count() interface{} {
	var count interface{}
	b.DB.Count(&count)

	return count
}
