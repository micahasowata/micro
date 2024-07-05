package memory

import (
	"context"
	"sync"

	"github.com/micahasowata/micro/metadata/internal/repository"
	"github.com/micahasowata/micro/metadata/pkg/model"
)

// Respository defines a memory movie metadata repository
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new memory respository
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrieves movies metadata by the movie id
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

// Put adds movie metadata for a given movie id
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata

	return nil
}
