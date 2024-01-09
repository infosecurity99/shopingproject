package structfortable

import "github.com/google/uuid"

type Users struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
