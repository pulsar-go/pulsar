package database

// Create the model on the database
func Create(model interface{}) *DB {
	return NewDB(db.Lib.Create(model))
}

// Save the model to the database
func Save(model interface{}) *DB {
	return NewDB(db.Lib.Save(model))
}

// Delete an existing model
func Delete(model interface{}, where ...interface{}) *DB {
	return db.Delete(model, where)
}

// Delete an existing model
func (d *DB) Delete(model interface{}, where ...interface{}) *DB {
	return NewDB(d.Lib.Delete(model, where))
}
