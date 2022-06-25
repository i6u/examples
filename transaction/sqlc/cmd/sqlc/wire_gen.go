// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/examples/transaction/sqlc/internal/biz"
	"github.com/go-kratos/examples/transaction/sqlc/internal/conf"
	"github.com/go-kratos/examples/transaction/sqlc/internal/data"
	"github.com/go-kratos/examples/transaction/sqlc/internal/server"
	"github.com/go-kratos/examples/transaction/sqlc/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db, cleanup := data.NewDB(confData, logger)
	dataData, cleanup2, err := data.NewData(db, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	transaction := data.NewTransaction(dataData)
	userUsecase := biz.NewUserUsecase(userRepo, transaction)
	transactionService := service.NewTransactionService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, transactionService)
	grpcServer := server.NewGRPCServer(confServer, logger, transactionService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
