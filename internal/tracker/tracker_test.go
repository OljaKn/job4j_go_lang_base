package tracker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/tracker"
)

func Test_Model(t *testing.T) {
	t.Parallel()
	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tr.UpdateItems(item)
		assert.ErrorIs(t, err, tracker.ErrNotFound)
	})

	t.Run("add fail", func(t *testing.T) {
		t.Parallel()

		tr := tracker.NewTracker()
		item := tracker.Item{
			ID:   "1",
			Name: "First Item",
		}
		itemNew := tracker.Item{
			ID:   "1",
			Name: "Second",
		}
		tr.AddItem(item)
		_, err := tr.AddItem(itemNew)
		assert.ErrorIs(t, err, tracker.ErrItemExist)
	})
}
