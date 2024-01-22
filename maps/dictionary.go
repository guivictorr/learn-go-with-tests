package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFoundWord = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExist = DictionaryErr("this word already exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFoundWord
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFoundWord:
		d[key] = value
	case nil:
		return ErrAlreadyExist
	default:
		return err
	}

	return nil
}
