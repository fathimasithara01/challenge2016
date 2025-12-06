package distributor

import "errors"

type Store struct {
	items map[string]*Distributor
}

func NewStore() *Store {
	return &Store{
		items: map[string]*Distributor{},
	}
}

func (s *Store) Add(name string, parent string) error {
	if name == "" {
		return errors.New("distributor name cannot be empty")
	}

	if parent != "" {
		if _, ok := s.items[parent]; !ok {
			return errors.New("parent distributor does not exist: " + parent)
		}
	}

	if _, exists := s.items[name]; exists {
		return nil
	}

	s.items[name] = &Distributor{Name: name, Parent: parent}
	return nil
}

func (s *Store) Get(name string) (*Distributor, bool) {
	d, ok := s.items[name]
	return d, ok
}

func (s *Store) ParentOf(name string) (string, bool) {
	d, ok := s.items[name]
	if !ok {
		return "", false
	}
	return d.Parent, true
}
