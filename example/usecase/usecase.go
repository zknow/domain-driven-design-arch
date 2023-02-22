package usecase

import (
	"context"

	"arch/domain"

	"errors"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type exampleUsecase struct {
	exampleRepo domain.ExampleRepository
}

// NewExampleRepoUsecase ...
func NewExampleRepoUsecase(exampleRepo domain.ExampleRepository) domain.ExampleUsecase {
	return &exampleUsecase{
		exampleRepo: exampleRepo,
	}
}

func (p *exampleUsecase) GetByID(ctx context.Context, id string) (*domain.Example, error) {
	data, err := p.exampleRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return data, nil
}

func (p *exampleUsecase) Store(ctx context.Context, d *domain.Example) error {
	if d.ID == "" {
		d.ID = uuid.Must(uuid.NewV4()).String()
	}
	if d.Status == "" {
		d.Status = "good"
	}
	if err := p.exampleRepo.Store(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (p *exampleUsecase) UpdateStatus(ctx context.Context, d *domain.Example) error {
	if d.Status == "" {
		err := errors.New("Status is blank")
		logrus.Error(err)
		return err
	}

	if err := p.exampleRepo.UpdateStatus(ctx, d); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
