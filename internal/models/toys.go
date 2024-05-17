package models

import (
	"errors"

	"github.com/vbrenister/green-toys/internal/validation"
)

var ErrRecordNotFound = errors.New("record not found")

var nextId = 0

func generateID() int {
	nextId++
	return nextId
}

type Toy struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Rented      bool    `json:"rented"`
}

func ValidateToy(v *validation.Validator, toy *Toy) {
	v.Check(toy.Title != "", "title", "must be provided")
	v.Check(toy.Description != "", "description", "must be provided")
	v.Check(toy.Price > 0, "price", "must be greater than zero")
	v.Check(toy.Category != "", "category", "must be provided")
}

type ToyModel interface {
	GetByID(id int) (*Toy, error)
	Create(toy *Toy) error
	GetAll() ([]*Toy, error)
}

type toyModel struct {
	toys map[int]*Toy
}

func (m *toyModel) GetByID(id int) (*Toy, error) {
	toy, ok := m.toys[id]
	if !ok {
		return nil, ErrRecordNotFound
	}

	return toy, nil
}

func (m *toyModel) Create(toy *Toy) error {
	id := generateID()
	toy.ID = id
	m.toys[id] = toy

	return nil
}

func (m *toyModel) GetAll() ([]*Toy, error) {
	toys := make([]*Toy, 0, len(m.toys))

	for _, toy := range m.toys {
		toys = append(toys, toy)
	}

	return toys, nil
}

func NewToyModel() ToyModel {
	firstID := generateID()
	return &toyModel{
		toys: map[int]*Toy{
			firstID: {
				ID:          firstID,
				Title:       "Toy",
				Description: "A toy",
				Price:       10.00,
				Category:    "Toy",
			},
		},
	}
}
