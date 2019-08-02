package maps

import "errors"

// Dictionary is custom type declaration
type Dictionary map[string]string

// ErrNotFound removes potential for clashing changes between code and tests
var ErrNotFound = errors.New("cannot find the word you're looking for")

// Search will comb through a map called dictionary for a word and
// return a definition.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
