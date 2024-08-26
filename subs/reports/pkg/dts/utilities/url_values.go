package utilities

import (
	"net/url"
	"slices"
	"strings"
)

// URLValues maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Values map
// are case-sensitive. URLValues differs from url.Values in that
// it allows adding values escaped AND unescaped.
type URLValues map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (uv URLValues) Get(key string) string {
	vs := uv[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (uv URLValues) Set(key, value string) {
	uv[key] = []string{value}
}

// SetEscape sets the key to the escaped value. It replaces
// any existing values.
func (uv URLValues) SetEscape(key, value string) {
	uv.Set(key, url.QueryEscape(value))
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (uv URLValues) Add(key, value string) {
	uv[key] = append(uv[key], value)
}

// AddEscape adds the escaped value to key. It appends to any
// existing values associated with key.
func (uv URLValues) AddEscape(key, value string) {
	uv.Add(key, url.QueryEscape(value))
}

// Del deletes the values associated with key.
func (uv URLValues) Del(key string) {
	delete(uv, key)
}

// Has checks whether a given key is set.
func (uv URLValues) Has(key string) bool {
	_, ok := uv[key]
	return ok
}

// Encode encodes the values into “URL encoded” form
// ("bar=baz&foo=quux") sorted by key.
func (uv URLValues) Encode() string {
	if len(uv) == 0 {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(uv))
	for k := range uv {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		vs := uv[k]
		keyEscaped := url.QueryEscape(k)
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			// we don't escape the value here because it's already escaped (if it was supposed to be)
			buf.WriteString(v)
		}
	}
	return buf.String()
}
