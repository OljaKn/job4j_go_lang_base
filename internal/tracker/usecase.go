package tracker

import (
	"github.com/google/uuid"
)

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()
	_, err := tracker.AddItem(Item{Name: name, ID: id})
	if err != nil {
		out.Out("failed add item")
	}
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.items {
		out.Out(item.toString())
	}
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter ID for update:")
	id := in.Get()
	_, err := tracker.FindById(id)
	if err != nil {
		out.Out("Not found Item")
		return
	}
	out.Out("Enter new name:")
	newName := in.Get()
	updateItem := Item{
		ID:   id,
		Name: newName,
	}
	if tracker.UpdateItems(updateItem) == nil {
		out.Out("Access")
	}
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter Id for delete:")
	id := in.Get()
	res := tracker.DeleteItems(id)
	if res == true {
		out.Out("Access")
	} else {
		out.Out("Not found Item")
	}
}

type FindByNameUsecase struct{}

func (u FindByNameUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("Enter string for found:")
	part := in.Get()
	res := tracker.FindByPartName(part)
	if len(res) == 0 {
		out.Out("Not found Item")
	}
	for _, item := range res {
		out.Out(item.toString())
	}
}
