package structfortable

import "github.com/google/uuid"

type OrderProducts struct {
	ID         uuid.UUID
	Order_id   uuid.UUID
	Product_id uuid.UUID
	Quantity   int
	Price      int
}
