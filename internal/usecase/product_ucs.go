package usecase

import (
	"github.com/Azmi117/API-USER2.git/internal/models"
	"github.com/Azmi117/API-USER2.git/internal/pkg/apperror"
	"github.com/Azmi117/API-USER2.git/internal/repository"
)

type ProductUsecase struct {
	repo *repository.ProductRepository
}

func NewProductUsecase(params *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repo: params,
	}
}

func (u *ProductUsecase) GetAll() ([]models.Product, error) {
	Products, err := u.repo.FindAll()

	if err != nil {
		return nil, apperror.Internal("Failed to get data!")
	}

	return Products, nil
}

func (u *ProductUsecase) GetById(id int) (models.Product, error) {
	// 1. Gunakan function repo untuk query, lalu tampung error jika ada
	Product, err := u.repo.FindById(id)

	// 2. Buat validasi jika gagal
	if err != nil {
		return models.Product{}, apperror.NotFound("Product with this id notfound!")
	}

	// 3. Return hasil
	return Product, nil
}

func (u *ProductUsecase) Create(body models.Product) (models.Product, error) {
	// 1. Cek apakah nama yang diinputkan sudah ada
	_, err := u.repo.FindByName(body.Name)

	if err == nil {
		return models.Product{}, apperror.BadRequest("Product is exist")
	}

	// 2. Buat validasi agar tidak ada input aneh
	if body.Name == "" {
		return models.Product{}, apperror.BadRequest("Invalid request body")
	}

	// 3. Gunakan function untuk query create
	res, err := u.repo.Create(body)

	// 4. Buat error jika gagal
	if err != nil {
		return models.Product{}, apperror.Internal("Failed create product")
	}

	// 5. Return value untuk response
	return res, nil
}

func (u *ProductUsecase) Update(id int, body models.Product) (models.Product, error) {
	// 1. Cek data existing
	existing, err := u.repo.FindById(id)

	if err != nil {
		return models.Product{}, apperror.NotFound("No product exist with this id")
	}

	// 2. buat Validasi untuk input aneh

	if body.Name == "" {
		body.Name = existing.Name
	}

	if body.Quantity < 0 {
		return models.Product{}, apperror.BadRequest("Quantity can't be minus")
	}

	// 3. Gunakan function repo untuk query ke db

	if err := u.repo.Update(body); err != nil {
		return models.Product{}, apperror.Internal("Failed update product")
	}

	// 4. Return hasil

	return body, nil
}

func (u *ProductUsecase) Delete(id int) error {
	// 1. Cek apakah data ada
	existing, err := u.repo.FindById(id)

	if err != nil {
		return apperror.NotFound("No product exist with this id")
	}

	// 2. Cek apakah product sudah pernah kena soft delete
	if existing.DeletedAt.Valid {
		return apperror.BadRequest("Product has been deleted")
	}

	// 3. Gunakan function repo untuk query
	if err := u.repo.Delete(existing); err != nil {
		return apperror.Internal("Failed deleted product")
	}

	// 4. Return hasil
	return nil
}
