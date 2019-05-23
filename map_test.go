package wxa

import (
	"errors"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	λ := func(e interface{}) (interface{}, error) {
		s, ok := e.(string)
		if !ok {
			return nil, errors.New("boom")
		}
		return strings.ToUpper(s), nil
	}
	f := Map(λ)
	t.Run("normal", func(t *testing.T) {
		result, err := f("a", "b")
		if err != nil {
			t.Fatal(err)
		}
		if result[0] != "A" {
			t.Errorf(`expected "A", got %q`, result[0])
		}
		if result[0] != "A" {
			t.Errorf(`expected "B", got %q`, result[1])
		}
	})
	t.Run("with errors", func(t *testing.T) {
		_, err := f(0)
		if err == nil {
			t.Fatal("expected an error")
		}
		if err.Error() != "boom" {
			t.Errorf(`expected "boom", got %q`, err.Error())
		}
	})
}
