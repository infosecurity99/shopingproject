package structfortable

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	ID        uuid.UUID
	Amount    int
	User_Id   uuid.UUID
	Create_At time.Time
}
