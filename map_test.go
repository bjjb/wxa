package Λ

import (
	"errors"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		λ := func(e interface{}) (interface{}, error) {
			return strings.ToUpper(e.(string)), nil
		}
		result, err := Map(λ, "a", "b")
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
		λ := func(e interface{}) (interface{}, error) {
			return nil, errors.New("boom")
		}
		_, err := Map(λ, "a", "b")
		if err == nil {
			t.Error("expected an error")
		}
		if err.Error() != "boom" {
			t.Errorf(`expected "boom", got %q`, err.Error())
		}
	})
}

func TestMustMap(t *testing.T) {
	λ := func(e interface{}) interface{} {
		return strings.ToUpper(e.(string))
	}
	result := MustMap(λ, "a", "b")
	if result[0] != "A" {
		t.Errorf(`expected "A", got %q`, result[0])
	}
	if result[0] != "A" {
		t.Errorf(`expected "B", got %q`, result[1])
	}
}
