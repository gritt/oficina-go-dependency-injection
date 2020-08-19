package core

type (
	TransactionSaver interface {
		Save(Transaction) error
	}

	Notifier interface {
		Notify(Transaction)
	}

	CreateUseCase struct {
		TransactionSaver
		Notifier
	}
)

func NewCreateUseCase(ts TransactionSaver, n Notifier) *CreateUseCase {
	return &CreateUseCase{ts, n}
}

func (c *CreateUseCase) Create(t Transaction) error {
	if err := t.Validate(); err != nil {
		return err
	}

	if err := c.Save(t); err != nil {
		return err
	}

	c.Notify(t)

	return nil
}
