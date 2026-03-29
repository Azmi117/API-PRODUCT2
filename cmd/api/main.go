package main

import (
	"fmt"
	"net/http"

	"github.com/Azmi117/API-USER2.git/internal/config"
	delivery "github.com/Azmi117/API-USER2.git/internal/delivery/http"
	"github.com/Azmi117/API-USER2.git/internal/repository"
	"github.com/Azmi117/API-USER2.git/internal/usecase"
)

func main() {
	db := config.ConnectDB()

	productRepo := repository.NewProductRepository(db)

	productUsecase := usecase.NewProductUsecase(productRepo)

	productHandler := delivery.NewProductHandler(productUsecase)

	mux := http.NewServeMux()

	delivery.MapRoutes(mux, productHandler)

	port := ":8080"

	fmt.Printf("Server running on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Failed run server: ", err)
	}
}
