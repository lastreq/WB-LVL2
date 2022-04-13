package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newDelimiter(t *testing.T) {
	input := []string{
		"aaa bb c",
		"aaabb c",
		"a aab bc",
	}
	expected := [][]string{
		{"aaa", "bb", "c"},
		{"aaabb", "c"},
		{"a", "aab", "bc"},
	}

	result := newDelimiter(input, " ")

	assert.True(t, reflect.DeepEqual(result, expected))
}

func Test_separatedStrings(t *testing.T) {
	input := [][]string{
		{"aaa", "bb", "c"},
		{"aaabb", "c"},
		{"ccccac"},
		{"a", "aab", "bc"},
		{"bbababa"},
	}
	expected := [][]string{
		{"aaa", "bb", "c"},
		{"aaabb", "c"},
		{"a", "aab", "bc"},
	}
	result := separatedStrings(input)
	assert.True(t, reflect.DeepEqual(expected, result))
}
