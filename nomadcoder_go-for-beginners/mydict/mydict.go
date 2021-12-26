package mydict

import "errors"

// Dictioinary type
type Dictioinary map[string]string

var errNotFound = errors.New("not Found")
var errWordExists = errors.New("that word already exists")

// Saerch for a word
func (d Dictioinary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictioinary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
