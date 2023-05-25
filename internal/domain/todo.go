package domain

import (
	"context"
	"time"
)

type TodoStatus uint8

const (
	TODO_ACTIVE TodoStatus = iota
	TODO_COMPLETED
)

type Todo struct {
	ID        string     `json:"id"`
	Task      string     `json:"task"`
	Status    TodoStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TodoRepository interface {
	GetAll(ctx context.Context) ([]Todo, error)
	// Find(ctx context.Context) (*Todo, error)
	Store(ctx context.Context, d *Todo) error
	// UpdateStatus(ctx context.Context, d *Todo) error
}

type TodoUsecase interface {
	GetAll(ctx context.Context) ([]Todo, error)
	// Find(ctx context.Context) (*Todo, error)
	Store(ctx context.Context, d *Todo) error
	// UpdateStatus(ctx context.Context, d *Todo) error
}
