package database

// IsNew checks if the model was already saved
func IsNew(model interface{}) bool {
	return db.Lib.NewRecord(model)
}
