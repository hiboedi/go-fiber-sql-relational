package migrations

import (
	"fmt"
	"go-fiber-relationaldb/database"
	"go-fiber-relationaldb/models"
)

func Migration() {
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
		&models.Post{},
		&models.Tag{},
	)

	if err != nil {
		fmt.Println("cant run migration")
	}

	fmt.Println("Migrated")
}
