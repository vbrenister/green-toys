package models

import (
	"errors"
	"fmt"
)

type Condition int

const (
	GoodCondition Condition = iota
	FairCondition
	PoorCondition
)

type Toy struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Rating      float64   `json:"rating"`
	Condition   Condition `json:"condition"`
	Category    string    `json:"category"`
}

func (c Condition) MarshalJSON() ([]byte, error) {
	var res string

	switch c {
	case GoodCondition:
		res = "good"
	case FairCondition:
		res = "fair"
	case PoorCondition:
		res = "poor"
	default:
		return nil, errors.New("invalid condition")
	}

	return []byte(fmt.Sprintf("\"%s\"", res)), nil
}
