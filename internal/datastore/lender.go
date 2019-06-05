package datastore

import (
	"time"

	"github.com/google/uuid"
)

type Lender struct {
	ID         int64
	Email      string
	GUID       uuid.UUID
	FirstName  string
	MiddleName string
	LastName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
