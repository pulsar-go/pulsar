package db

// First finds the first record with the given condition
func (b *DB) First(out interface{}, where ...interface{}) *DB {
	return b.clone(b.DB.First(out, where...))
}

// Last finds the last record with the given condition
func (b *DB) Last(out interface{}, where ...interface{}) *DB {
	return b.clone(b.DB.Last(out, where...))
}

// All finds all records of a given model
func (b *DB) All(models interface{}, where ...interface{}) *DB {
	return b.clone(b.DB.Find(models, where...))
}

// Where adds a condition to the query statement
func (b *DB) Where(query interface{}, args ...interface{}) *DB {
	query = enhanceQuery(query)
	return b.clone(b.DB.Where(query, args...))
}

// WhereNot adds a condition to the query statement
func (b *DB) WhereNot(query interface{}, args ...interface{}) *DB {
	return b.clone(b.Not(query, args...))
}

// OrWhere adds an or filter to the query
func (b *DB) OrWhere(query interface{}, args ...interface{}) *DB {
	return b.clone(b.Or(query, args...))
}
