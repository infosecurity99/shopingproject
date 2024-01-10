package structfortable

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	ID        uuid.UUID `json:"id"`
	Amount    int       `json:"amount"`
	UserId   uuid.UUID `json:"userid"`
	CreateAt time.Time `json:"createat"`
}
