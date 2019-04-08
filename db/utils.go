package db

import "strings"

// IsNew checks if the model was already saved
func (b *DB) IsNew(model interface{}) bool {
	return b.NewRecord(model)
}

// enhanceQuery params
func enhanceQuery(query interface{}) interface{} {
	if s, ok := query.(string); ok && !strings.Contains(s, "?") {
		return query.(string) + " = ?"
	}

	return query
}
