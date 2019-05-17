package lernen

import (
	"github.com/jtelgen/weblib/api"
)

// GelerntVon view gelernt-von
type GelerntVon struct {
	api.View
}

// GelerntVonRow row from view gelernt-von
type GelerntVonRow struct {
	ID  string   `json:"id"`
	Key []string `json:"key"`
	Rev string   `json:"value"`
}

// AllDocs returns all Docs
func (v GelerntVon) AllDocs(rows *[]GelerntVonRow) error {
	return v.View.AllDocs(rows)
}

// DocsByKey returns all Docs matching the given key
func (v GelerntVon) DocsByKey(key string, rows *[]GelerntVonRow) error {
	return v.View.DocsByKey(key, rows)
}
