package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_LruCache_Methods(t *testing.T) {
	t.Parallel()

	t.Run("Add and get nodes", func(t *testing.T) {
		cache := base.NewLruCache(3)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		assert.Equal(t, "world", *res)
	})

	t.Run("Node is not exists", func(t *testing.T) {
		cache := base.NewLruCache(3)
		cache.Put("hello", "world")
		cache.Put("one", "two")
		cache.Put("two", "three")
		res := cache.Get("three")
		assert.Nil(t, res)
	})

	t.Run("Cache size = 1", func(t *testing.T) {
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		res := cache.Get("hello")
		assert.Equal(t, "world", *res)
	})

	t.Run("Put exsisting node", func(t *testing.T) {
		cache := base.NewLruCache(3)
		cache.Put("hello", "world")
		cache.Put("one", "two")
		cache.Put("hello", "1111")
		res := cache.Get("hello")
		assert.NotEqual(t, "world", *res)
	})

	t.Run("Capacity > size", func(t *testing.T) {
		cache := base.NewLruCache(2)
		cache.Put("hello", "world")
		cache.Put("one", "two")
		cache.Put("key", "1111")
		res := cache.Get("hello")
		assert.Nil(t, res)
	})

	t.Run("Cache size = 1, put exists node", func(t *testing.T) {
		cache := base.NewLruCache(1)
		cache.Put("hello", "world")
		cache.Put("two", "value")
		res := cache.Get("hello")
		res2 := cache.Get("two")
		assert.Nil(t, res)
		assert.Equal(t, "value", *res2)
	})
}
