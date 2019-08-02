package maps

// Search will comb through a map called dictionary for a word and
// return a definition.
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
