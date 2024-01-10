package structfortable

import "github.com/google/uuid"

type Users struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}
