package main

import (
	"GoMicBase/app/account/conf"
	"GoMicBase/app/account/database"
	"GoMicBase/app/account/server"
	"GoMicBase/app/account/service"
	"GoMicBase/pkg/registry"
	"GoMicBase/pkg/zlog"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type accountApp struct {
	config    *conf.AccountConfig
	server    *grpc.Server
	registery *registry.ConsulRegistery
}

func initApp(path, name string) *accountApp {
	var err error
	accountConfig, err := conf.NewAccountConfig(path, name)
	if err != nil {
		zlog.Panicln(err.Error())
	}

	db, err := database.NewMysqlDB(&accountConfig.MysqlConfig)
	if err != nil {
		zlog.Panicln(err.Error())
	}
	accountService := service.NewAccountService(db)
	consulRegistery, err := registry.NewConsulRegistery(accountConfig.ConsulConfig.Host, accountConfig.ConsulConfig.Port)
	if err != nil {
		zlog.Panicln(err.Error())
	}
	grpcServer := server.NewGrpcServer(accountService)

	app := &accountApp{
		config:    accountConfig,
		server:    grpcServer,
		registery: consulRegistery,
	}

	return app
}

func (app *accountApp) startServ() {
	dsn := fmt.Sprintf("%s:%d", app.config.AccountServConfig.Host, app.config.AccountServConfig.Port)
	listen, err := net.Listen("tcp", dsn)
	if err != nil {
		zlog.Panicln(err.Error())
	}
	zlog.Infof("Start Account GRPC Service on: %s", dsn)

	if err := app.registery.NewClient(); err != nil {
		zlog.Panicln(err.Error())
	}

	accountServConf := app.config.AccountServConfig
	if err := app.registery.RegisterGrpcServ(accountServConf.Host, int(accountServConf.Port), accountServConf.Name, accountServConf.Id); err != nil {
		zlog.Panicln(err.Error())
	}

	go func() {
		if err := app.server.Serve(listen); err != nil {
			zlog.Panicln(err.Error())
		}
	}()
}

func (app *accountApp) stopServ() {
	if err := app.registery.Deregister(app.config.AccountServConfig.Id); err != nil {
		zlog.Infoln(err.Error())
	}

	app.server.GracefulStop()
}
