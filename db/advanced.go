package db

// Limit adds a limit to the query
func (b *DB) Limit(limit interface{}) *DB {
	return b.clone(b.DB.Limit(limit))
}

// Skip adds an offset to the query
func (b *DB) Skip(skip interface{}) *DB {
	return b.clone(b.Offset(skip))
}

// Count 's the database records in the given table
func (b *DB) Count() interface{} {
	var count interface{}
	b.DB.Count(&count)

	return count
}

// Table specify the table you would like to run db operations
func (b *DB) Table(name string) *DB {
	return b.clone(b.DB.Table(name))
}

// Select specifies field you want to query from the database
func (b *DB) Select(query interface{}, args ...interface{}) *DB {
	return b.clone(b.DB.Select(query, args...))
}

// Group specify the group method on the find
func (b *DB) Group(query string) *DB {
	return b.clone(b.DB.Group(query))
}

// Having specify HAVING conditions for GROUP BY
func (b *DB) Having(query string) *DB {
	return b.clone(b.DB.Having(query))
}

// Scan scan value to a struct
func (b *DB) Scan(dest interface{}) *DB {
	return b.clone(b.DB.Scan(dest))
}
