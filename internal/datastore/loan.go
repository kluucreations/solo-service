package datastore

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	ID          int64
	Name        string
	Description string
	BorrowerID  int64
	GUID        uuid.UUID
	Amount      float64
	StartsAt    time.Time
	EndsAt      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d datastore) GetLoansByAmount(page, perPage int64, asc bool) (loans []Loan, err error) {
	limit, offset := pagination(page, perPage)
	rows, err := d.Query(`
	SELECT
		l.id, l.name, l.description, l.borrower_id, l.guid,
		l.amount, l.starts_at, l.ends_at, l.created_at, l.updated_at
	FROM
		loans l
	LEFT JOIN
		lender_loans ll on l.id = ll.loan_id
	WHERE 
		ll.loan_id IS NULL
	ORDER BY
		l.amount $1
	LIMIT $2 OFFSET $3;`,
		orderBy(asc), limit, offset,
	)
	defer rows.Close()
	if err != nil {
		return
	}
	return
}
func (d datastore) GetLoansByEndDate(page, perPage int64, asc bool) (loans []Loan, err error) {
	limit, offset := pagination(page, perPage)
	rows, err := d.Query(`
	SELECT
		l.id, l.name, l.description, l.borrower_id, l.guid,
		l.amount, l.starts_at, l.ends_at, l.created_at, l.updated_at
	FROM
		loans l
	LEFT JOIN
		lender_loans ll on l.id = ll.loan_id
	WHERE 
		ll.loan_id IS NULL
	ORDER BY
		l.ends_at $1
	LIMIT $2 OFFSET $3;`,
		orderBy(asc), limit, offset,
	)
	defer rows.Close()
	if err != nil {
		return
	}
	return
}
func (d datastore) GetLoansByScore(page, perPage int64, asc bool) (loans []Loan, err error) {
	limit, offset := pagination(page, perPage)
	rows, err := d.Query(`
	SELECT
		l.id, l.name, l.description, l.borrower_id, l.guid,
		l.amount, l.starts_at, l.ends_at, l.created_at, l.updated_at
	FROM
		loans l
	LEFT JOIN
		lender_loans ll on l.id = ll.loan_id
	WHERE 
		ll.loan_id IS NULL
	ORDER BY
		b.score $1
	LIMIT $2 OFFSET $3;`,
		orderBy(asc), limit, offset,
	)
	defer rows.Close()
	if err != nil {
		return
	}
	return
}
