package database

import (
	"dumbflix/models"
	"dumbflix/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Film{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
