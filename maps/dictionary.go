package maps

import "errors"

// Dictionary is custom type declaration
type Dictionary map[string]string

// Search will comb through a map called dictionary for a word and
// return a definition.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("cannot find the word you're looking for")
	}
	return definition, nil
}
