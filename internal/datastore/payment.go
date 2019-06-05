package datastore

type Payment struct {
}

func (d datastore) GetPayments(loanID int64) (payments []Payment, err error) {
	rows, err := d.Query(`
	SELECT
		id, 
	FROM
		payments
	WHERE
		loan_id
	`)
	defer rows.Close()
	if err != nil {
		return
	}
	return
}
