package main

import (
	"github.com/gritt/oficina-go-dependency-injection/core"
	"github.com/gritt/oficina-go-dependency-injection/details/cli"
	"github.com/gritt/oficina-go-dependency-injection/details/db"
	"github.com/gritt/oficina-go-dependency-injection/details/mailer"
)

/*
 * Validar e Salvar uma Transaction
 * - identificar e lidar com erros nesse fluxo [foi]
 * - automatizar injeção de dependências [ongoing]
 */
func main() {
	app := initCLI()
	app.Run()

	// api := initAPI()
	// api.Serve()
}

func initCLI() *cli.CLI {
	// repository [detail]
	dbCfg := db.NewConfig()
	repo := db.NewRepository(dbCfg) // inject db config

	// mail notifier [detail]
	mCfg := mailer.NewConfig()
	noti := mailer.NewEmailSender(mCfg) // inject mail config

	// use case [core]
	useCase := core.NewCreateUseCase(repo, noti) // inject repo as TransactionSaver, mailer as Notifier

	// run as a CLI app [detail]
	app := cli.NewCLI(useCase) // inject userCase as Creator

	return app
}
