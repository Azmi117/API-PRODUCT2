package main

import (
	"fmt"

	"github.com/Azmi117/API-USER2.git/internal/config"
	"github.com/Azmi117/API-USER2.git/internal/models"
)

func main() {
	// 1. inisialisasi variable untuk function connect db
	db := config.ConnectDB()

	err := db.AutoMigrate(&models.Product{})

	if err != nil {
		fmt.Println("Failed migration to db : ", err)
		return
	}

	fmt.Println("Migration Success!")
}
