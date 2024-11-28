package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ivegotanidea/golang-gorm-postgres/initializers"
	. "github.com/ivegotanidea/golang-gorm-postgres/models"
	"gorm.io/gorm"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
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

	if existingUser.Error != nil {
		panic(existingUser.Error)
	}

	if existingUser.RowsAffected != 0 {
		fmt.Println("🧘‍♂️He Who Remains is there ⏳")
	}

	if err := db.Where("role = ?", "owner").FirstOrCreate(&owner).Error; err != nil {
		panic(err)
	}
}

func Init() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err := initializers.DB.AutoMigrate(
		&City{},
		&Ethnos{},
		&BodyType{},
		&BodyArt{},
		&ProfileBodyArt{},
		&HairColor{},
		&IntimateHairCut{},
		&Payment{},
		&Photo{},
		&Profile{},
		&ProfileOption{},
		&ProfileRating{},
		&ProfileTag{},
		&RatedProfileTag{},
		&RatedUserTag{},
		&Service{},
		&User{},
		&UserRating{},
		&UserTag{})

	// Auto-migrate the User model
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	CreateOwnerUser(initializers.DB)

	if err := initializers.DB.Exec("CREATE UNIQUE INDEX unique_owner ON users (tier) WHERE tier = 'owner'").Error; err != nil {
		log.Fatalf("Failed to create unique index: %v", err)
	}

	fmt.Println("👍 Migration complete")
}

func main() {
	Init()
}
