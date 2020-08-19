package cli

import (
	"errors"
	"fmt"

	"github.com/gritt/oficina-go-dependency-injection/core"
)

type (
	TransactionCreator interface {
		Create(t core.Transaction) error
	}

	CLI struct {
		TransactionCreator
	}
)

func NewCLI(creator TransactionCreator) *CLI {
	return &CLI{creator}
}

func (c *CLI) Run() {
	// TODO: read os.Args
	trs := core.Transaction{
		Name:   "test",
		Amount: 10,
	}

	if err := c.Create(trs); err != nil {
		fmt.Printf("%T \n", err)

		dbErr := &core.DatabaseError{}
		if errors.As(err, &dbErr) {
			fmt.Println(err)
		}

		tsErr := &core.TransactionError{}
		if errors.As(err, &tsErr) {
			fmt.Println(err)
		}
	}

	fmt.Println("=== exit ===")
}
