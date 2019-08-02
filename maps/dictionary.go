package maps

// Dictionary is custom type declaration
type Dictionary map[string]string

// Search will comb through a map called dictionary for a word and
// return a definition.
func (d Dictionary) Search(word string) string {
	return d[word]
}
