package cb

import (
	"github.com/gocql/gocql"
)

// Write write
type Write struct{}

// NewWrite new write
func NewWrite() *Write {
	return &Write{}
}

// List set list
func (w *Write) List(query string, args []interface{}) error {
	b := session.NewBatch(gocql.LoggedBatch)
	for _, v := range args {
		b.Query(query, v.([]interface{})...)
	}
	return session.ExecuteBatch(b)
}

// Item set item
func (w *Write) Item(query string, args ...interface{}) error {
	q := session.Query(query, args...)
	return q.Exec()
}

// ItemNamed set item named
func (w *Write) ItemNamed(query string, args interface{}) error {
	s := NewNamed().Set(query)
	m := NewNamed().Map(args)
	t := NewNamed().Interface(s, m)

	q := session.Query(query, t...)
	return q.Exec()
}
