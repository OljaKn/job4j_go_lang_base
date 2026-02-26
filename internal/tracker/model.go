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

func (t *Tracker) AddItem(item Item) (Item, error) {
	index, ok := t.indexOf(item.ID)
	if ok {
		return t.items[index], ErrItemExist
	}
	t.items = append(t.items, item)
	return item, nil
}

func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}

func (t *Tracker) FindById(id string) (*Item, error) {
	index, ok := t.indexOf(id)
	if !ok {
		return nil, ErrIdNotFound
	}
	return &t.items[index], nil
}

func (t *Tracker) DeleteItems(id string) bool {
	index, ok := t.indexOf(id)
	if !ok {
		return false
	}
	t.items = append(t.items[:index], t.items[index+1:]...)
	return true
}

func (t *Tracker) UpdateItems(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}
	t.items[index].Name = item.Name
	return nil
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

func (t *Tracker) indexOf(id string) (int, bool) {
	for i := 0; i < len(t.items); i++ {
		if t.items[i].ID == id {
			return i, true
		}
	}
	return -1, false
}
