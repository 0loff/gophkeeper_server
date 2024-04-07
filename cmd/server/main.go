package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/0loff/gophkeeper_server/internal/app"
	"github.com/0loff/gophkeeper_server/internal/logger"
	"github.com/0loff/gophkeeper_server/internal/server"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var ErrorGracefulStop = errors.New("connection have been closed")

func main() {
	Run()
}

func Run() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	a := app.NewApp()
	s := server.NewServer(a.UserUC, a.DataUC)

	s.Init()

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		if err := logger.Initialize("info"); err != nil {
			log.Fatal(err)
		}

		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			log.Fatal(err)
		}

		logger.Sugar.Infoln("The Grpc server is running on port :3200")

		return s.Srv.Serve(listen)
	})
	g.Go(func() error {
		<-gCtx.Done()
		s.Srv.GracefulStop()
		return ErrorGracefulStop
	})

	if err := g.Wait(); err != nil {
		logger.Log.Error("Server have been stopped with error", zap.Error(err))
	}
}
