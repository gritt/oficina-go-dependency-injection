package core

import "fmt"

type (
	TransactionError struct {
		Operation   string
		Message     string
		Transaction Transaction
	}

	DatabaseError struct {
		Query string
		Err   error
	}
)

func (e *TransactionError) Error() string {
	return fmt.Sprintf(
		`transaction error: %s failed with msg "%s": "%s", "%d"`,
		e.Operation, e.Message, e.Transaction.Name, e.Transaction.Amount,
	)
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf(`database error: query "%s" failed with msg: "%s"`, e.Query, e.Err.Error())
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}
