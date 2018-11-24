package cb

import (
	"regexp"

	"github.com/fatih/structs"
)

// Named named
type Named struct{}

// NewNamed new named
func NewNamed() *Named {
	return &Named{}
}

// Set query string to set
func (n *Named) Set(query string) []string {
	reg := regexp.MustCompile(`:(.*?)(,|\))`)
	return reg.FindAllString(query, -1)
}

// Map args interface to map
func (n *Named) Map(args interface{}) map[string]interface{} {
	m := structs.New(args)
	m.TagName = "json"
	return m.Map()
}

// Interface set & map to interface
func (n *Named) Interface(s []string, m map[string]interface{}) []interface{} {
	t := []interface{}{}
	for _, v := range s {
		key := v[1 : len(v)-1]
		if value, ok := m[key]; ok {
			t = append(t, value)
		}
	}
	return t
}
