package mydict

import "errors"

// Dictioinary type
type Dictioinary map[string]string

var errNotFound = errors.New("not Found")

// Saerch for a word
func (d Dictioinary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
