package domain

import "context"

// Example ...
type Example struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// ExampleRepository ...
type ExampleRepository interface {
	GetByID(ctx context.Context, id string) (*Example, error)
	Store(ctx context.Context, d *Example) error
	UpdateStatus(ctx context.Context, d *Example) error
}

// ExampleUsecase ..
type ExampleUsecase interface {
	GetByID(ctx context.Context, id string) (*Example, error)
	Store(ctx context.Context, d *Example) error
	UpdateStatus(ctx context.Context, d *Example) error
}
