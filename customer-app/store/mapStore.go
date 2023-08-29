package store

import (
	"customerapp/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (m MapStore) Create(c domain.Customer) error {
	if _, ok := m.store[c.ID]; ok {
		return domain.ErrCustomerIdAlreadyExist
	}
	m.store[c.ID] = c
	return nil
}

func (m MapStore) Update(id string, c domain.Customer) error {
	if _, ok := m.store[id]; !ok {
		return domain.ErrCustomerNotFound
	}
	m.store[id] = c
	return nil
}

func (m MapStore) Delete(id string) error {
	if _, ok := m.store[id]; !ok {
		return domain.ErrCustomerNotFound
	}
	delete(m.store, id)
	return nil
}

func (m MapStore) GetById(id string) (domain.Customer, error) {
	if _, ok := m.store[id]; !ok {
		return domain.Customer{}, domain.ErrCustomerNotFound
	}
	return m.store[id], nil
}

func (m MapStore) GetAll() ([]domain.Customer, error) {
	if len(m.store) == 0 {
		return []domain.Customer{}, domain.ErrNoCustomerExists
	}
	var customers []domain.Customer
	for _, v := range m.store {
		customers = append(customers, v)
	}
	return customers, nil
}
