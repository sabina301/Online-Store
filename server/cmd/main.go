package main

import (
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
		DBName:   "online-store",
		SSLMode:  "disable",
	}
	db, err := repository.NewDatabase(dbConf)

	if err != nil {
		log.Fatalf("Error: unable to connect to database")
	}

	userRep := repository.NewUserRepository(db.GetDB())
	productRep := repository.NewProductRepository(db.GetDB())
	userService := service.NewUserService(userRep)
	productService := service.NewProductService(productRep)
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	router.InitRouter(userHandler, productHandler)
}
