package maps

import "testing"

func assertMessage(x testing.TB, want string, got string) {
	x.Helper()
	if got != want {
		x.Errorf("expected, %s, got %s", want, got)
	}
}
func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error")
	}
}
func assertDefinition(t testing.TB, dictionary Dictionary, key string, want string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertMessage(t, want, got)
}

func assertErrorMessage(t testing.TB, want error, got error) {
	t.Helper()
	if want != got {
		t.Errorf("expected, %s, got %s", want, got)
	}
}

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is test string"}
	t.Run("word found", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is test string"
		assertMessage(t, want, got)
	})
	t.Run("word not found", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err)
		assertErrorMessage(t, err, ErrKeyNotFound)
	})

}

func TestAdd(t *testing.T) {

	assertKeyExist := func(t testing.TB, got error, want error) {
		t.Helper()
		if got != want {
			t.Errorf("expected, %s, got %s", want, got)
		}
	}

	t.Run("add item", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		dictionary.Add(key, "new item added")
		want := "new item added"
		assertDefinition(t, dictionary, key, want)
	})
	t.Run("add existing item", func(t *testing.T) {
		key := "test"
		value := "new item added"

		dictionary := Dictionary{key: value}
		err := dictionary.Add("test", "new other item added")
		assertKeyExist(t, err, ErrKeyFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update found item", func(t *testing.T) {
		key := "test"
		value := "update me"
		dictionary := Dictionary{key: value}
		newValue := "updated me"

		dictionary.Update(key, newValue)

		assertDefinition(t, dictionary, key, newValue)
	})
	t.Run("not found key update", func(t *testing.T) {
		key := "test2"
		value := "update me"
		dictionary := Dictionary{"test": value}
		newValue := "updated me"

		err := dictionary.Update(key, newValue)
		assertError(t, err)
		assertErrorMessage(t, ErrUpdateKeyNotFound, err)
	})

}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "testing delete"}
	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	assertError(t, err)
	assertErrorMessage(t, ErrKeyNotFound, err)

}
