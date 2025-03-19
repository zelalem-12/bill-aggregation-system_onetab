package seeder

import (
	"log"

	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/adapter/persistence/postgres/model"
	"github.com/zelalem-12/bill-aggregation-system_onetab/user-service/internal/util"
	"gorm.io/gorm"
)

func SeedDemoData(db *gorm.DB) {

	rawPassword := "SecurePass123!"

	hashedPassword, err := util.HashPassword(rawPassword)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	var existingUser model.User
	if err := db.Where("email = ?", "john.doe@example.com").First(&existingUser).Error; err == nil {
		log.Println("User with the given email already exists.")
		return
	}

	user := model.User{
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		Password:       string(hashedPassword),
		IsVerified:     true,
		ProfilePicture: "https://example.com/profiles/john.jpg",
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to seed demo user: %v", err)
	} else {
		log.Println("Demo user seeded successfully.")
		log.Printf("Use this email to login: %s", user.Email)
		log.Printf("Use this password to login: %s", rawPassword)
	}
}
