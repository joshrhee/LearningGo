package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("word is not found")
var errWordExists = errors.New("that word already exists")
var errCantUpdate = errors.New("can't update non existing word")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, isExist := d[word]
	if isExist {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

//Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
