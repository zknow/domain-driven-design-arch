package repository

import (
	"context"
	"database/sql"

	"arch/domain"
)

type exampleRepository struct {
	db *sql.DB
}

// NewExampleRepository ...
func NewExampleRepository(db *sql.DB) domain.ExampleRepository {
	return &exampleRepository{db}
}

func (p *exampleRepository) GetByID(ctx context.Context, id string) (*domain.Example, error) {
	d := &domain.Example{}
	return d, nil
}

func (p *exampleRepository) Store(ctx context.Context, d *domain.Example) error {
	return nil
}

func (p *exampleRepository) UpdateStatus(ctx context.Context, d *domain.Example) error {
	return nil
}
