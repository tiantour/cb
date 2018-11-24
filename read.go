package cb

import (
	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
)

// Read read
type Read struct{}

// NewRead new read
func NewRead() *Read {
	return &Read{}
}

// List get list output
func (r *Read) List(dest interface{}, query string, args ...interface{}) *gocql.Iter {
	q := session.Query(query, args...)
	return q.Iter()
}

// Item get item output struct
func (r *Read) Item(query string, args ...interface{}) *gocql.Query {
	return session.Query(query, args...)
}

// ListStruct get list output struct
func (r *Read) ListStruct(dest interface{}, query string, args ...interface{}) error {
	q := session.Query(query, args...)
	m, err := q.Iter().SliceMap()
	if err != nil {
		return err
	}
	return mapstructure.Decode(m, dest)
}

// ItemStruct get item output struct
func (r *Read) ItemStruct(dest interface{}, query string, args ...interface{}) error {
	m := map[string]interface{}{}
	q := session.Query(query, args...)
	err := q.MapScan(m)
	if err != nil {
		return err
	}
	return mapstructure.Decode(m, dest)
}
