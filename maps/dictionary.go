package maps

import "errors"

// Dictionary is custom type declaration
type Dictionary map[string]string

// ErrNotFound removes potential for clashing changes between code and tests
var (
	ErrNotFound   = errors.New("cannot find the word you're looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search will comb through a map called dictionary for a word and
// return a definition.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add will append the new word and definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
