package model

import (
	"errors"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Item

func (t *Todos) Add(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	items := *t

	if index <= 0 || index > len(items) {
		return errors.New("Invalid index")
	}

	items[index-1].Done = true
	items[index-1].CompletedAt = time.Now()

	return nil
}

func (t *Todos) Delete(index int) error {
	items := *t

	if index <= 0 || index > len(items) {
		return errors.New("Invalid index")
	}

	*t = append(items[:index-1], items[index:]...)

	return nil
}
