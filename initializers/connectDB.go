package initializers

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	duration, err := time.ParseDuration(config.DBQueriesSlowThreshold)

	if err != nil {
		log.Fatal(err)
	}

	// Create a custom logger
	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Output logs to the console
		logger.Config{
			SlowThreshold: duration,
			LogLevel:      logger.LogLevel(config.DBLogLevel), // Log level: Info logs all queries
			Colorful:      true,                               // Enable color output
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

func CreateOwnerUser(db *gorm.DB) {

	owner := User{
		ID:             uuid.Max,
		Name:           "He Who Remains",
		Phone:          "77778889900",
		TelegramUserId: 6794234746,
		Password:       "h5sh3d", // Ensure this is hashed
		Avatar:         "https://akm-img-a-in.tosshub.com/indiatoday/images/story/202311/tom-hiddleston-in-a-still-from-loki-2-27480244-16x9_0.jpg",
		Verified:       true,
		HasProfile:     false,
		Tier:           "owner",
	}

	existingUser := db.Where("id =?", owner.ID).First(&owner)

	if existingUser.RowsAffected != 0 {
		fmt.Println("🧘‍♂️He Who Remains is there ⏳")
		return
	}

	if err := db.Where("role = ?", "owner").FirstOrCreate(&owner).Error; err != nil {
		panic(err)
	}

	if err := DB.Exec("CREATE UNIQUE INDEX unique_owner ON users (tier) WHERE tier = 'owner'").Error; err != nil {
		log.Fatalf("Failed to create unique index: %v", err)
	}
}

func Migrate() {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	log.Printf("Automigrating BASE models...")

	err := DB.AutoMigrate(
		// no needs
		&City{},
		&Ethnos{},
		&BodyType{},
		&BodyArt{},
		&HairColor{},
		&IntimateHairCut{},
		&UserTag{},
		&ProfileTag{},
		&User{})

	if err != nil {
		log.Fatalf("Failed to auto-migrate BASE models: %v", err)
	}

	log.Printf("Automigrating T-1 models...")
	err = DB.AutoMigrate(
		&Profile{},         // needs User
		&Payment{},         // needs User
		&Photo{},           // needs Profile
		&RatedProfileTag{}, // needs ProfileTag
		&RatedUserTag{},    // needs UserTag
	)

	if err != nil {
		log.Fatalf("Failed to auto-migrate T-1 models: %v", err)
	}

	log.Printf("Automigrating T-2 models...")
	err = DB.AutoMigrate(
		&Service{},        // needs User, Profile
		&ProfileBodyArt{}, // needs Profile, BodyArt
		&ProfileOption{},  // needs Profile, ProfileTag
	)

	if err != nil {
		log.Fatalf("Failed to auto-migrate T-2 models: %v", err)
	}

	log.Printf("Automigrating T-3 models...")
	err = DB.AutoMigrate(
		&ProfileRating{}, // needs User, Profile, Service, RatedProfileTag
		&UserRating{},    // needs User, Service, RatedProfileTag
	)

	if err != nil {
		log.Fatalf("Failed to auto-migrate T-3 models: %v", err)
	}

	fmt.Println("Creating owner users...")
	CreateOwnerUser(DB)
	fmt.Println("Creating owner users... OK")

	fmt.Println("👍 Migration complete")
}
