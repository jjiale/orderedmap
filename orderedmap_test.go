package orderedmap_test

import (
	//"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderedMap(t *testing.T) {
	m := NewOrderedMap()
	assert.IsType(t, &OrderedMap{}, m)
}

func TestGet(t *testing.T) {
	t.Run("ReturnsNotOKIfStringKeyDoesntExist", func(t *testing.T) {
		m := NewOrderedMap()
		_, ok := m.Get("foo")
		assert.False(t, ok)
	})

	t.Run("ReturnsNotOKIfNonStringKeyDoesntExist", func(t *testing.T) {
		m := NewOrderedMap()
		_, ok := m.Get(123)
		assert.False(t, ok)
	})

	t.Run("ReturnsOKIfKeyExists", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "bar")
		_, ok := m.Get("foo")
		assert.True(t, ok)
	})

	t.Run("ReturnsValueForKey", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "bar")
		value, _ := m.Get("foo")
		assert.Equal(t, "bar", value)
	})

	t.Run("ReturnsDynamicValueForKey", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "baz")
		value, _ := m.Get("foo")
		assert.Equal(t, "baz", value)
	})

	t.Run("KeyDoesntExistOnNonEmptyMap", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "baz")
		_, ok := m.Get("bar")
		assert.False(t, ok)
	})

	t.Run("ValueForKeyDoesntExistOnNonEmptyMap", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "baz")
		value, _ := m.Get("bar")
		assert.Nil(t, value)
	})

}

func TestSet(t *testing.T) {
	t.Run("ReturnsTrueIfStringKeyIsNew", func(t *testing.T) {
		m := NewOrderedMap()
		ok := m.Set("foo", "bar")
		assert.True(t, ok)
	})

	t.Run("ReturnsTrueIfNonStringKeyIsNew", func(t *testing.T) {
		m := NewOrderedMap()
		ok := m.Set(123, "bar")
		assert.True(t, ok)
	})

	t.Run("ValueCanBeNonString", func(t *testing.T) {
		m := NewOrderedMap()
		ok := m.Set(123, true)
		assert.True(t, ok)
	})

	t.Run("SetThreeDifferentKeys", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "bar")
		m.Set("baz", "qux")
		ok := m.Set("quux", "corge")
		assert.True(t, ok)
	})

}

func TestAdd(t *testing.T) {
	t.Run("Returns3IfAdd2", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", 1)
		m.Add("foo", 2)
		got, _ := m.Get("foo")
		assert.Equal(t, got, 3)
	})

	t.Run("ReturnsFalseIfValueIsNotInt", func(t *testing.T) {
		m := NewOrderedMap()
		m.Set("foo", "bar")
		ok := m.Add("foo", 1)
		assert.False(t, ok)
	})
}
