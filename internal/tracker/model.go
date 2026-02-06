package tracker

import (
	"fmt"
	"strings"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) toString() string {
	return fmt.Sprintf("%s\t%s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
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

func (t *Tracker) FindByName(name string) *Item {
	for i := 0; i < len(t.items); i++ {
		if t.items[i].Name == name {
			return &t.items[i]
		}
	}
	return nil
}

func (t *Tracker) DeleteItems(name string) bool {
	for i := 0; i < len(t.items); i++ {
		if t.items[i].Name == name {
			t.items = append(t.items[:i], t.items[i+1:]...)
			return true
		}
	}
	return false
}

func (t *Tracker) FindByPartName(part string) []*Item {
	res := []*Item{}
	for i := 0; i < len(t.items); i++ {
		if strings.Contains(t.items[i].Name, part) {
			res = append(res, &t.items[i])
		}
	}
	return res
}
