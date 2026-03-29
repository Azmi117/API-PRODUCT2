package repository

import (
	"github.com/Azmi117/API-USER2.git/internal/models"
	"gorm.io/gorm"
)

// 1. Buat struct yg return db

type ProductRepository struct {
	db *gorm.DB
}

// 2. Buat constructor dari struct di atas
func NewProductRepository(params *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: params,
	}
}

// 3. Bikin method untuk query (GetAll, GetById, Create, dll)
func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var Products []models.Product
	err := r.db.Find(&Products).Error
	return Products, err
}

func (r *ProductRepository) FindById(id int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) FindByName(name string) (models.Product, error) {
	var product models.Product
	err := r.db.Where("name = ?", name).First(&product).Error
	return product, err
}

func (r *ProductRepository) Create(body models.Product) (models.Product, error) {
	err := r.db.Create(&body).Error
	return body, err
}

func (r *ProductRepository) Update(body models.Product) error {
	return r.db.Save(&body).Error
}

func (r *ProductRepository) Delete(body models.Product) error {
	return r.db.Delete(&body).Error
}
