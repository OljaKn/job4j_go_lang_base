package tracker

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(string)
}

type Store interface {
	Create(ctx context.Context, item Item) error
	List(ctx context.Context) ([]Item, error)
	Get(ctx context.Context, id string) (Item, error)
	Update(ctx context.Context, item Item) error
	Delete(ctx context.Context, id string) error
	FindByPartName(ctx context.Context, part string) ([]Item, error)
}

type Usecase interface {
	Done(ctx context.Context, in Input, out Output, store Store) error
}

type AddUsecase struct{}

func (u AddUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("enter name:")
	name := in.Get()
	id := uuid.New().String()

	if err := store.Create(
		ctx,
		Item{ID: id, Name: name},
	); err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}
	return nil
}

type GetUsecase struct{}

func (u GetUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	items, err := store.List(ctx)
	if err != nil {
		return fmt.Errorf("failed to get items: %w", err)
	}
	for _, item := range items {
		out.Out(item.ID + " " + item.Name)
	}
	return nil
}

type UpdateUsecase struct{}

func (u UpdateUsecase) Done(
	ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("enter ID for update:")
	id := in.Get()
	_, err := store.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("not found Item with id: %w", err)
	}
	out.Out("Enter new name:")
	newName := in.Get()
	updateItem := Item{
		ID:   id,
		Name: newName,
	}
	if store.Update(ctx, updateItem) == nil {
		out.Out("Access")
	}
	return nil
}

type DeleteUsecase struct{}

func (u DeleteUsecase) Done(ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("Enter Id for delete:")
	id := in.Get()
	err := store.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("not found Item for delete with id: %w", err)
	}
	out.Out("Access")
	return nil
}

type FindByNameUsecase struct{}

func (u FindByNameUsecase) Done(ctx context.Context,
	in Input,
	out Output,
	store Store,
) error {
	out.Out("Enter string for found:")
	part := in.Get()
	items, err := store.FindByPartName(ctx, part)
	if err != nil {
		out.Out("Not found Item")
	}
	for _, item := range items {
		out.Out(item.ID + " " + item.Name)
	}
	return nil
}
