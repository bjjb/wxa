package wxa

import (
	"errors"
	"testing"
)

func TestFind(t *testing.T) {
	λ := func(e interface{}) (bool, error) {
		i, ok := e.(int)
		if !ok {
			return false, errors.New("boom")
		}
		return i%2 == 0, nil
	}
	f := Find(λ)
	t.Run("normal", func(t *testing.T) {
		r, err := f(1, 2, 3, 4)
		if err != nil {
			t.Fatal(err)
		}
		if r != 2 {
			t.Errorf("expected 2, got %d", r.(int))
		}
	})
	t.Run("no results", func(t *testing.T) {
		r, err := f(1, 3, 5, 7, 9)
		if err != nil {
			t.Fatal("didn't expect an error")
		}
		if r != nil {
			t.Fatal("expected result to be nil")
		}
	})
	t.Run("error", func(t *testing.T) {
		_, err := f("not an int")
		if err == nil {
			t.Errorf("expected an error")
		}
		if err.Error() != "boom" {
			t.Errorf(`expected the error to be "boom", it was %s`, err.Error())
		}
	})
}
