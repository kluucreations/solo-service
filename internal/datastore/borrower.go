package datastore

import (
	"time"

	"github.com/google/uuid"
)

type Borrower struct {
	ID         int64
	Email      string
	GUID       uuid.UUID
	FirstName  string
	MiddleName string
	LastName   string
	Score      Score
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
