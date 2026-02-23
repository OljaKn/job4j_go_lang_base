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

func (t *Tracker) FindById(id string) *Item {
	index := t.IndexOf(id)
	if index == -1 {
		return nil
	}
	return &t.items[index]
}

func (t *Tracker) DeleteItems(id string) bool {
	index := t.IndexOf(id)
	if index == -1 {
		return false
	}
	t.items = append(t.items[:index], t.items[index+1:]...)
	return true
}

func (t *Tracker) UpdateItems(id string, newName string) bool {
	index := t.IndexOf(id)
	if index == -1 {
		return false
	}
	t.items[index].Name = newName
	return true
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

func (t *Tracker) IndexOf(id string) int {
	for i := 0; i < len(t.items); i++ {
		if t.items[i].ID == id {
			return i
		}
	}
	return -1
}
