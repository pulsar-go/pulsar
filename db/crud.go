package db

// Model specify the model you would like to run db operations
func (b *DB) Model(value interface{}) *DB {
	return b.clone(b.DB.Model(value))
}

// Create insert the value into database
func (b *DB) Create(value interface{}) *DB {
	return b.clone(b.DB.Create(value))
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (b *DB) Save(value interface{}) *DB {
	return b.clone(b.DB.Save(value))
}

// Update update attributes with callbacks
func (b *DB) Update(attrs ...interface{}) *DB {
	return b.clone(b.DB.Update(attrs...))
}

// Updates update attributes with callbacks
func (b *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) *DB {
	return b.clone(b.DB.Updates(values, ignoreProtectedAttrs...))
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
func (b *DB) Delete(value interface{}, where ...interface{}) *DB {
	return b.clone(b.DB.Delete(value, where...))
}

// Unscoped return all record including deleted record
func (b *DB) Unscoped() *DB {
	return b.clone(b.DB.Unscoped())
}
