package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tracker struct {
	items []Item
}

type Item struct {
	ID   string
	Name string
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func Test_NewTracker(t *testing.T) {
	t.Parallel()
	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t,
			[]Item{item},
			tracker.GetItems(),
		)
	})
}
