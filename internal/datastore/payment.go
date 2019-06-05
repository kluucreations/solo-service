package datastore

import (
	"time"

	"github.com/google/uuid"
)

type PaymentStatus int

const (
	PaymentStatusInit PaymentStatus = iota
	PaymentStatusPending
	PaymentStatusSuccess
	PaymentStatusError
	PaymentStatusDelinquent
)

type Payment struct {
	ID         int64
	GUID       uuid.UUID
	Amount     float64
	Status     PaymentStatus
	PaymentAt  time.Time
	BorrowerID int64
	LenderID   int64
	Breakdown  interface{} // Not Yet Implemented
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (d datastore) GetPayments(loanID int64) (payments []Payment, err error) {
	rows, err := d.Query(`
	SELECT
		id, guid, loan_id, iso_currency_code, amount, status, borrower_id, 
		created_at, updated_at
	FROM
		payments
	WHERE
		loan_id = $1
	ORDER BY
		id desc;
	`, loanID)
	defer rows.Close()
	if err != nil {
		return
	}
	return
}
