package main

import (
	"flag"
	server "github.com/deuuus/bmstu-rsoi"
	"github.com/deuuus/bmstu-rsoi/pkg/handler"
	"github.com/deuuus/bmstu-rsoi/pkg/repository"
	"github.com/deuuus/bmstu-rsoi/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	configName := flag.String("config", "config", "Path to configuration file")

	if err := initConfig(*configName); err != nil {
		logrus.Fatalf("error while config initializition: %s", err.Error())
	}

	logrus.Info(configName)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	logrus.Info(viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.username"))

	if err != nil {
		logrus.Fatalf("error while db initializition: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server.Run(handlers.InitRoutes())
}

func initConfig(configName string) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}
