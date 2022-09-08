package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/quankori/go-clear-architecture/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	echo   *echo.Echo
	cfg    *config.Configuration
	db     *mongo.Database
	logger *logrus.Logger
	ready  chan bool
}

func NewServer(cfg *config.Configuration, db *mongo.Database, logger *logrus.Logger, ready chan bool) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger, ready: ready}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         ":" + "8200",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		s.logger.Logf(logrus.InfoLevel, "Server is listening on PORT: %s", "8200")
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalln("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	if s.ready != nil {
		s.ready <- true
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.logger.Fatalln("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
