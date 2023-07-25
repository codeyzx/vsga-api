package migration

import (
	"fmt"
	"log"
	"vsga-api/database"
	"vsga-api/model/entity"
)

func RunMigration() {

	err := database.DB.AutoMigrate(&entity.Employe{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
