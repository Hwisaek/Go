package mydict

import "errors"

// Dictioinary type
type Dictioinary map[string]string

var (
	errNotFound   = errors.New("not Found")
	errCantUpdate = errors.New("can't update non-existing word")
	errCantDelete = errors.New("can't delete non-existing word")
	errWordExists = errors.New("that word already exists")
)

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

// Update a word to the dictionary
func (d Dictioinary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word to the dictionary
func (d Dictioinary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
