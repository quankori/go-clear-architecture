package main

import (
	"log"

	"github.com/quankori/go-clear-architecture/config"
	"github.com/quankori/go-clear-architecture/drivers/mongo"
	"github.com/quankori/go-clear-architecture/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {

	log.Println("Starting api server")
	// Initialize config
	cfg := config.NewConfig()
	db := mongo.ConnectMongoDB(cfg.DatabaseConnectionURL, "test")
	s := server.NewServer(cfg, db, logrus.New(), nil)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
