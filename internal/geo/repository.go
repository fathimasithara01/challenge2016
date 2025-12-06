package geo

import "sync"

type Repository struct {
	mu      sync.RWMutex
	regions map[string]Region
}

func NewRepository() *Repository {
	return &Repository{
		regions: make(map[string]Region),
	}
}

func (r *Repository) Add(region Region) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := region.City + "|" + region.State + "|" + region.Country
	r.regions[key] = region
}
