package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"server/handler"
	"server/repository"
	"server/router"
	"server/service"
)

func main() {
	dbConf := repository.DatabaseConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres15",
		Password: "password",
		DBName:   "online-store-db",
		SSLMode:  "disable",
	}
	db, err := repository.NewDatabase(dbConf)

	if err != nil {
		log.Fatalf("Error: unable to connect to database")
	}

	rep := repository.NewRepository(db.GetDB())
	serv := service.NewService(rep)
	hand := handler.NewHandler(serv)

	r := router.InitRouter(hand)

	err = router.Start(r, "0.0.0.0:8080")
	if err != nil {
		logrus.Fatal("Cant run")
	}
}
