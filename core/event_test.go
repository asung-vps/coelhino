package core

import (
	"testing"
)

func TestRemoveWhiteSpace(t *testing.T) {
	t.Run("Simple test", func(t *testing.T) {
		target := "white space"
		expected := "whitespace"
		actual := removeWhiteSpace(target)
		if expected != actual {
			msg := "Target mismatch.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})

	t.Run("Escaped space", func(t *testing.T) {
		target := "white\\ space"
		expected := "white space"
		actual := removeWhiteSpace(target)
		if expected != actual {
			msg := "Target mismatch.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})

	t.Run("Multiple whitespace", func(t *testing.T) {
		target := "Please   remove\n\tthese\t\t multiple    whitespaces"
		expected := "Pleaseremovethesemultiplewhitespaces"
		actual := removeWhiteSpace(target)
		if expected != actual {
			msg := "Target mismatch.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})

	t.Run("Out of bounds edge case", func(t *testing.T) {
		target := " whitespace"
		expected := "whitespace"
		actual := removeWhiteSpace(target)
		if expected != actual {
			msg := "Target mismatch.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})
}

func TestParseToEventRunStatement(t *testing.T) {
	stmt := "run:transaction"
	event, _ := ParseToEvent(stmt)

	t.Run("Test command extraction", func(t *testing.T) {
		expected := "run"
		actual := event.GetCommand()
		if expected != actual {
			msg := "Command mismatch.\nExpected: %v\nGot: %v"
			t.Errorf(msg, expected, actual)
		}
	})

	t.Run("Test argument extraction", func(t *testing.T) {
		expectedList := []string{"transaction"}
		actualList := event.GetArgs()
		for i := 0; i < len(actualList); i++ {
			expected := expectedList[i]
			actual := actualList[i]
			if expected != actual {
				msg := "Arg mismatch at index %v.\nExpected: %v\nGot: %v"
				t.Errorf(msg, i, expected, actual)
			}
		}
	})
}
