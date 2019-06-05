package datastore

import "database/sql"

type datastore struct {
	sql.DB
}

type Datastore interface {
	GetPayments(loanID int64) ([]Payment, error)
}

func NewClient() Datastore {
	return datastore{}
}
