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
		assert.Equal(t, res, "world")
	})
}
