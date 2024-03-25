package app

import (
	"log"

	"github.com/0loff/gophkeeper_server/config"
	"github.com/0loff/gophkeeper_server/internal/data"
	dataRepository "github.com/0loff/gophkeeper_server/internal/data/repository/postgres"
	dataUsecases "github.com/0loff/gophkeeper_server/internal/data/usecases"
	"github.com/0loff/gophkeeper_server/internal/db/postgres"
	"github.com/0loff/gophkeeper_server/internal/user"
	userRepository "github.com/0loff/gophkeeper_server/internal/user/repository/postgres"
	userUsecases "github.com/0loff/gophkeeper_server/internal/user/usecases"
)

type App struct {
	Cfg config.Config

	UserUC user.UserProcessor
	DataUC data.DataProcessor
}

func NewApp() *App {
	app := &App{Cfg: config.NewConfigBuilder()}

	db, err := postgres.InitDB(app.Cfg.DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	ur := userRepository.NewUserRepository(db)
	app.UserUC = userUsecases.NewUserUseCases(ur)

	dr := dataRepository.NewDataRepository(db)
	app.DataUC = dataUsecases.NewDataUseCases(dr)

	return app
}
