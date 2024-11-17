package initializers

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	// Create a custom logger
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Output logs to the console
		logger.Config{
			SlowThreshold: time.Millisecond, // Threshold for slow query logging
			LogLevel:      logger.Info,      // Log level: Info logs all queries
			Colorful:      true,             // Enable color output
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger},
	)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
