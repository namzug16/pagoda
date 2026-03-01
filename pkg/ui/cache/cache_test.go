package cache

import (
	"testing"

	. "github.com/namzug16/gotags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache_GetSet(t *testing.T) {
	key := "test"
	assert.Nil(t, Get(key))

	node := Div("hello")
	Set(key, node)

	got := Get(key)
	require.NotNil(t, got)

	// Check it was converted to a Raw component.
	_, ok := got.(Raw)
	require.True(t, ok)

	// Both nodes should render the same string.
	assert.Equal(t, node.String(), got.String())
}

func TestCache_SetIfNotExists(t *testing.T) {
	key := "test2"
	called := 0
	callback := func() HTML {
		called++
		return Div("hello")
	}

	assertRender := func(n HTML) {
		assert.Equal(t, `<div>hello</div>`, n.String())
	}

	got := SetIfNotExists(key, callback)
	assert.Equal(t, 1, called)
	require.NotNil(t, got)
	assertRender(got)

	got = SetIfNotExists(key, callback)
	assert.Equal(t, 1, called)
	require.NotNil(t, got)
	assertRender(got)
}
