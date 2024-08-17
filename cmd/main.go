package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"

	"ticket-booking-app/cmd/api"
	"ticket-booking-app/cmd/config"
	"ticket-booking-app/internal/db/connection"
	"ticket-booking-app/internal/i18n"
)

var (
	once      sync.Once
	conn      *gorm.DB
	serveConf config.ServeConfig
)

func init() {
	once.Do(func() {
		conn = connection.PostgresSQLConnection(connection.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			AppName:  os.Getenv("APP_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
			Timezone: os.Getenv("DB_TIMEZONE"),
		})

	})

	serveConf = config.ServeConfig{
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
	}

	i18n.InitBundle("./internal/i18n/languages/")
}

func main() {

	app := fiber.New(config.FiberConfig)

	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02T15:04:05.000Z",
		TimeZone:   "Europe/Istanbul",
	}))

	api.InitializeRouter(app, conn)

	// Start listening on port 8000
	go func() {
		if err := app.Listen(":" + serveConf.Port); err != nil {
			panic(err)
		}
	}()

	// Graceful shutdown
	err := GracefulShutdown(app, 5*time.Second)
	if err != nil {
		log.Error("Graceful shutdown error", err)
	}
}

func GracefulShutdown(app *fiber.App, timeout time.Duration) error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	sig := <-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db, err := conn.DB()
	if err != nil {
		return err
	}

	if err := shutdownDatabase(ctx, db); err != nil {
		return err
	}

	if err := app.Shutdown(); err != nil {
		return err
	}

	log.Infof("Signal received: %v", sig)
	return nil
}

func shutdownDatabase(ctx context.Context, db *sql.DB) error {
	ch := make(chan error, 1)
	go func() {
		ch <- db.Close()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		if err != nil {
			log.Error("Database close error", err)
			return err
		}
		return nil
	}
}
