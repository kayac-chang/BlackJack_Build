package frame

import "strings"

// A MetaData is a map of key-values pairs
type MetaData map[string][]string

// Add appends the given values to the key.
func (m MetaData) Add(key string, values ...string) {
	if len(values) == 0 {
		return
	}
	key = strings.ToLower(key)
	m[key] = append(m[key], values...)
}

// Del removes the values associated with the key.
func (m MetaData) Del(key string) {
	key = strings.ToLower(key)
	delete(m, key)
}

// Get returns the first value associated with the key.
func (m MetaData) Get(key string) string {
	key = strings.ToLower(key)
	if s, ok := m[key]; ok && len(s) > 0 {
		return s[0]
	}
	return ""
}

// Values returns all values associated with the key.
func (m MetaData) Values(key string) []string {
	key = strings.ToLower(key)
	return m[key]
}

// Set replaces the existing values associated with the key.
// If the length of values is 0, it removes all existing values.
func (m MetaData) Set(key string, values ...string) {
	key = strings.ToLower(key)
	m[key] = values
}
