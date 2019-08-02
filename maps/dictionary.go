package maps

// Dictionary is custom type declaration
type Dictionary map[string]string

// DictionaryErr establishes type that accesses the error interface
type DictionaryErr string

// ErrNotFound defines error message for Search
const (
	ErrNotFound         = DictionaryErr("cannot find the word you're looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

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

// Update takes word and definition and will update the definition of any
// existing word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete takes a word and removes it from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
