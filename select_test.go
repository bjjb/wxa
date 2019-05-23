package wxa

import (
	"errors"
	"testing"
)

func TestSelect(t *testing.T) {
	λ := func(e interface{}) (bool, error) {
		i, ok := e.(int)
		if !ok {
			return false, errors.New("boom")
		}
		return i%2 == 0, nil
	}
	f := Select(λ)
	t.Run("with some results", func(t *testing.T) {
		r, err := f(1, 2, 3, 4)
		if err != nil {
			t.Fatal(err)
		}
		if r == nil {
			t.Fatal("expected some results")
		}
		if len(r) != 2 {
			t.Errorf(`expected 2 results, got %d`, len(r))
		}
		if r[0] != 2 {
			t.Errorf(`expected 2, got %d`, r[0])
		}
		if r[1] != 4 {
			t.Errorf(`expected 4, got %d`, r[0])
		}
	})
	t.Run("with no results", func(t *testing.T) {
		r, err := f(1, 3, 5, 7, 9)
		if err != nil {
			t.Fatal("didn't expect an error")
		}
		if r == nil {
			t.Fatal("didn't expect results to be nil")
		}
		if len(r) != 0 {
			t.Fatalf("expected length of results to be 0, got %d", len(r))
		}
	})
	t.Run("with errors", func(t *testing.T) {
		_, err := f("not an int")
		if err == nil {
			t.Fatal("expected an error")
		}
		if err.Error() != "boom" {
			t.Errorf(`expected "boom", got %q`, err.Error())
		}
	})
}
