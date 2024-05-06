package maps

type Dictionary map[string]string

type DictionaryErr string

func (x DictionaryErr) Error() string {
	return string(x)
}

const (
	ErrKeyNotFound       = DictionaryErr("could not find key value")
	ErrKeyFound          = DictionaryErr("key already exists")
	ErrUpdateKeyNotFound = DictionaryErr("update key not found")
)

func (x Dictionary) Search(key string) (string, error) {
	value, found := x[key]
	if !found {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (x Dictionary) Add(key string, value string) error {
	_, err := x.Search(key)
	switch err {
	case ErrKeyNotFound:
		x[key] = value
	case nil:
		return ErrKeyFound
	default:
		return err
	}
	return nil
}

func (x Dictionary) Update(key string, value string) error {
	_, err := x.Search(key)
	switch err {
	case nil:
		x[key] = value
	case ErrKeyNotFound:
		return ErrUpdateKeyNotFound
	default:
		return err
	}
	return nil
}

func (x Dictionary) Delete(key string) {
	delete(x, key)
}
