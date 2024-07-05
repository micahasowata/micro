package metadata

import (
	"context"
	"errors"

	"github.com/micahasowata/micro/metadata/internal/repository"
	"github.com/micahasowata/micro/metadata/pkg/model"
)

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller defines a metadata service controller
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo: repo}
}

// Get returns movie metadata by id
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, repository.ErrNotFound
		}

		return nil, err
	}

	return res, nil
}
