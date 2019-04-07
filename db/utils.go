package db

// IsNew checks if the model was already saved
func (b *DB) IsNew(model interface{}) bool {
	return b.NewRecord(model)
}
