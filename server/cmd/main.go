package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"server/handler"
	"server/repository"
	"server/router"
	"server/service"
	"syscall"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("Error with configuration: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error with env file")
	}

	dbConf := repository.DatabaseConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("db_password"),
		DBName:   viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := repository.NewDatabase(dbConf)

	if err != nil {
		log.Fatalf("Error: unable to connect to database")
	}
	defer db.CloseDB()
	rep := repository.NewRepository(db.GetDB())
	serv := service.NewService(rep)
	hand := handler.NewHandler(serv)

	r := router.InitRouter(hand)

	srv := new(router.Server)
	_, err = serv.CreateAdmin()
	if err != nil {
		logrus.Println("Cant create admin")
	}
	go func() {
		err = srv.Start(r, viper.GetString("port"))
		if err != nil {
			logrus.Fatal("Cant run")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	err = srv.Finish(context.Background())
	if err != nil {
		logrus.Error(err.Error())
	}
	err = db.CloseDB()
	if err != nil {
		logrus.Error(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configuration")
	viper.SetConfigName("configuration")
	return viper.ReadInConfig()
}
