package structfortable

import "github.com/google/uuid"

type Products struct {
	ID uuid.UUID  `json:"id"`
	Price  int   `json:"price"`
	Name string  `json:"name"`
}