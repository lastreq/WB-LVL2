package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWithoutEscapeExamples(t *testing.T) {
	require.Equal(t, "aaaabccddddde", unpackString("a4bc2d5e"), "check a4bc2d5e")
	require.Equal(t, "abcd", unpackString("abcd"), "check abcd")
	require.Equal(t, "не корректная строка", unpackString("45"), "check 45")
}

func TestWithEscapeExamples(t *testing.T) {
	require.Equal(t, `qwe45`, unpackString(`qwe\4\5`), `check qwe\4\5`)
	require.Equal(t, `qwe44444`, unpackString(`qwe\45`), `check qwe\45`)
	require.Equal(t, `qwe\\\\\`, unpackString(`qwe\\5`), `check qwe\\5`)
}
