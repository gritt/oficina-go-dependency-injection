// +build wireinject

package main

// var createUseCaseDeps = wire.NewSet(
// 	mailer.NewConfig,
// 	mailer.NewEmailSender,
// 	wire.Bind(new(core.Notifier), new(*mailer.EmailSender)),
//
// 	db.NewConfig,
// 	db.NewRepository,
// 	wire.Bind(new(core.TransactionSaver), new(*db.Repository)),
// )

// func initCLI() *cli.CLI {
// 	panic(wire.Build(
// 		mailer.NewConfig,
// 		mailer.NewEmailSender,
// 		wire.Bind(new(core.Notifier), new(*mailer.EmailSender)),
//
// 		db.NewConfig,
// 		db.NewRepository,
// 		wire.Bind(new(core.TransactionSaver), new(*db.Repository)),
//
// 		core.NewCreateUseCase,
// 		wire.Bind(new(cli.TransactionCreator), new(*core.CreateUseCase)),
//
// 		cli.NewCLI,
// 	))
// }

// func initAPI() *rest.API {
// 	panic(wire.Build(
// 		mailer.NewConfig,
// 		mailer.NewEmailSender,
// 		wire.Bind(new(core.Notifier), new(*mailer.EmailSender)),
//
// 		db.NewConfig,
// 		db.NewRepository,
// 		wire.Bind(new(core.TransactionSaver), new(*db.Repository)),
//
// 		core.NewCreateUseCase,
// 		wire.Bind(new(rest.TransactionCreator), new(*core.CreateUseCase)),
//
// 		rest.NewAPI,
// 	))
// }
