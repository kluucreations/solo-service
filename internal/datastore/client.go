package datastore

import "database/sql"

type datastore struct {
	sql.DB
}

type Datastore interface {
	GetPayments(loanID int64) ([]Payment, error)
	GetLoansByAmount(page, perPage int64, asc bool) ([]Loan, error)
	GetLoansByEndDate(page, perPage int64, asc bool) ([]Loan, error)
	GetLoansByScore(page, perPage int64, asc bool) ([]Loan, error)
}

func NewClient() Datastore {
	return datastore{}
}

func orderBy(asc bool) string {
	if asc {
		return "ASC"
	}
	return "DESC"
}

func pagination(page, perPage int64) (limit, offset int64) {
	return perPage, (page - 1) * page
}
