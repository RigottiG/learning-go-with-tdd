package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordAlreadyExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist  = DictionaryErr("cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// It's a method that takes a dictionary and a word and returns a definition and an
// error.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Checking if the word exists in the dictionary. If it does, it returns an error.
// If it doesn't, it adds the word to the dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}

	return nil
}

// It's checking if the word exists in the dictionary. If it does, it returns an error.
// If it doesn't, it adds the word to the dictionary.
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

// It's a method that takes a dictionary and a word and deletes the word from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
