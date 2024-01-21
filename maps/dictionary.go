package maps

import "errors"

type Dictionary map[string]string

var ErrNotFoundWord = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFoundWord
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
