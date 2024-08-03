package utils

import (
	"fmt"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
	"math/rand/v2"
)

func CleanupTestUsers(db *gorm.DB) {
	db.Where("name LIKE ?", "test%").Delete(&models.User{})
}

func DropAllTables(db *gorm.DB) {

	tables, err := db.Migrator().GetTables()

	if err != nil {
		message, _ := fmt.Printf("failed to get tables: %s", err)
		panic(message)
	}

	for _, table := range tables {
		err = db.Migrator().DropTable(table)
		if err != nil {
			message, _ := fmt.Printf("failed to drop table: %s", err)
			panic(message)
		}
	}
}

func GenerateRandomPhoneNumber(r *rand.Rand, length int) string {

	if length <= 0 {
		length = 11 // Default length
	}

	minLen := int64(1)

	for i := 1; i < length; i++ {
		minLen *= 10
	}

	maxLen := minLen * 10
	return fmt.Sprintf("%0*d", length, r.Int64N(maxLen-minLen)+minLen)
}
