package structfortable

import "github.com/google/uuid"

type OrderProducts struct {
	ID         uuid.UUID
	OrderId   uuid.UUID
	ProductId uuid.UUID
	Quantity   int
	Price      int
}
