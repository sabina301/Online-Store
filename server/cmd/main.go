package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"server/handler"
	"server/repository"
	"server/router"
	"server/service"
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

	err = srv.Start(r, viper.GetString("port"))
	if err != nil {
		logrus.Fatal("Cant run")
	}

}

func initConfig() error {
	viper.AddConfigPath("configuration")
	viper.SetConfigName("configuration")
	return viper.ReadInConfig()
}
