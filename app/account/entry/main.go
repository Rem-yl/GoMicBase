package main

import (
	"GoMicBase/pkg/zlog"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "config", "dev", "config path, eg: -config dev")
}

func main() {
	flag.Parse()
	app := initApp("../", flagconf)
	app.startServ()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	zlog.Infoln("Received interrupt signal, shutting down gracefully...")
	app.stopServ()
	zlog.Infoln("Grpc server shut down gracefully")
}
