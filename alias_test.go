package di

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAliasMap(t *testing.T) {
	a := AliasMap{
		"a": "b",
		"c": "d",
	}
	b := a.Copy()

	require.Equal(t, a, b)
	require.Equal(t, "b", a.Get("a"))
	require.Equal(t, "e", a.Get("e"))
}
