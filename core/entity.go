package core

type Transaction struct {
	Amount int
	Name   string
}

func (t Transaction) Validate() error {
	if t.Amount <= 0 {
		return &TransactionError{"Transaction.Validate", "invalid amount", t}
	}
	if t.Name == "" {
		return &TransactionError{"Transaction.Validate", "invalid name", t}
	}
	return nil
}
