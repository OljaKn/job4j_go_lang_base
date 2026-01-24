package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_NewTracker(t *testing.T) {
	t.Parallel()
	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]base.Item{item},
			tracker.GetItems(),
		)
	})

	t.Run("items not equal", func(t *testing.T) {
		t.Parallel()

		tracker := base.NewTracker()
		item := base.Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		copy := tracker.GetItems()
		copy[0].Name = "Second Item"
		copy2 := tracker.GetItems()

		assert.Equal(t,
			"First Item",
			copy2[0].Name,
		)
		assert.Equal(t, "Second Item", copy[0].Name)
	})
}
