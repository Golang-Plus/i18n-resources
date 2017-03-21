package overwriting

import (
	"github.com/golang-plus/errors"
	"github.com/golang-plus/i18n"
)

// Values represents the data of resource values.
// Langauge Code : Value
type Values map[string]string

// Resource represents a multi-language resource.
type Resource struct {
	Language *i18n.Language
	Values   Values
}

// Get returns the resource value by given key.
func (r *Resource) Get(key string) string {
	if len(key) > 0 {
		if v, ok := r.Values[key]; ok {
			return v
		}
	}

	return ""
}

// Set sets the resource value.
func (r *Resource) Set(key, value string) {
	r.Values[key] = value
}

// NewResource returns a new resource.
func NewResource(language *i18n.Language, values Values) (*Resource, error) {
	if language == nil {
		return nil, errors.New("language of resource cannot be nil")
	}

	vals := values
	if vals == nil {
		vals = make(Values)
	}

	return &Resource{
		Language: language,
		Values:   vals,
	}, nil
}

// MustNewResource is like as NewResource but panic if error happens.
func MustNewResource(language *i18n.Language, values Values) *Resource {
	r, err := NewResource(language, values)
	if err != nil {
		panic(err)
	}

	return r
}
