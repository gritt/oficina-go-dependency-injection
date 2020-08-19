package mailer

import (
	"fmt"

	"github.com/gritt/oficina-go-dependency-injection/core"
)

type EmailSender struct {
	cfg *Config
}

func NewEmailSender(cfg *Config) *EmailSender {
	return &EmailSender{cfg}
}

func (m EmailSender) Notify(t core.Transaction) {
	fmt.Printf("sending email to: %s .... \n", m.cfg.Server)
}
