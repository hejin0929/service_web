// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"world/internal/biz"
	"world/internal/conf"
	"world/internal/data"
	"world/internal/server"
	"world/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewMysql(confData, logger)
	client := data.NewRedis(confData, logger)
	rockscacheClient := data.NewRocksCache(confData, logger, client)
	dataData, cleanup, err := data.NewData(confData, logger, db, client, rockscacheClient)
	if err != nil {
		return nil, nil, err
	}
	countryRepo := data.NewCountryRepo(dataData, logger)
	countryUse := biz.NewCountryUse(countryRepo, logger)
	countryService := service.NewCountryService(countryUse)
	grpcServer := server.NewGRPCServer(confServer, countryService, logger)
	httpServer := server.NewHTTPServer(confServer, countryService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
