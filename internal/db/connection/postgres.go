package connection

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ticket-booking-app/internal/db/models"
)

var once sync.Once

func PostgresSQLConnection(config DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s application_name='%s' sslmode=%s timezone=%s",
		config.Host,
		config.Username,
		config.Password,
		config.DBName,
		config.Port,
		config.AppName,
		config.SSLMode,
		config.Timezone,
	)

	connection, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil
	}

	// Migrate the database
	migration(connection)

	return connection
}

func migration(connection *gorm.DB) {
	// Auto migrate
	once.Do(func() {

		log.Info("Migrating the database...")

		err := connection.AutoMigrate(
			&models.Ticket{},
			&models.Event{},
			&models.User{},
		)
		if err != nil {
			log.Error("Error migrating the database: ", err)
		} else {
			log.Info("Database migration is successful.")
		}
	})
}
